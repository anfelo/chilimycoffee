package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "layout", data)
}

func main() {
	templates := make(map[string]*template.Template)
	templates["home"] = template.Must(template.ParseFiles("html/home.html", "html/layout.html"))
	templates["login"] = template.Must(template.ParseFiles("html/login.html", "html/layout.html"))
	templates["guide"] = template.Must(template.ParseFiles("html/guide.html", "html/layout.html"))

	e := echo.New()
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:CSRF",
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Static("/static", "public")

	e.GET("/", Home)
	e.GET("/login", Login)
	e.POST("/login", LoginSubmit)

	e.GET("/guides/:guide_slug", Guide)
	e.GET("/guides/:guide_slug/:part_slug", GuidePart)

	e.Logger.Fatal(e.Start(":3000"))
}

func Home(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	// Check if user is authenticated
	// if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
	// 	return c.Redirect(http.StatusFound, "/login")
	// }

	data := make(map[string]interface{})
	data["FlashMessages"] = sess.Flashes("flash")
	sess.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "home", data)
}

type ChapterPartConfig struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Chapter string `toml:"chapter"`
}

type GuideConfig struct {
	Title        string              `toml:"title"`
	Slug         string              `toml:"slug"`
	Description  string              `toml:"description"`
	Chapters     []string            `toml:"chapters"`
	ChapterParts []ChapterPartConfig `toml:"chapter_parts"`
}

func Guide(c echo.Context) error {
	guideSlug := c.Param("guide_slug")

	var guideConf GuideConfig
	_, err := toml.DecodeFile(fmt.Sprintf("./public/guides/%s/content.toml", guideSlug), &guideConf)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Get first part slug. Test this scenario
	if len(guideConf.ChapterParts) == 0 {
		return c.Redirect(http.StatusNotFound, "/")
	}

	firstPart := guideConf.ChapterParts[0]
	return c.Redirect(http.StatusFound, fmt.Sprintf("/guides/%s/%s", guideSlug, firstPart.Slug))
}

func GuidePart(c echo.Context) error {
	type partData struct {
		Title string
		Path  string
	}
	type guideData struct {
		Title    string
		Chapters map[string][]partData
		Content  template.HTML
	}

	guideSlug := c.Param("guide_slug")
	partSlug := c.Param("part_slug")

    var guideTitle string
	var guideConf GuideConfig
	_, err := toml.DecodeFile(fmt.Sprintf("./public/guides/%s/content.toml", guideSlug), &guideConf)
	if err != nil {
        // TODO: Redirect to the home page maybe, with a flash message
	}

    guideTitle = guideConf.Title
	partContent, err := os.ReadFile(fmt.Sprintf("./public/guides/%s/%s.html", guideSlug, partSlug))
	if err != nil {
        guideTitle = "Not Found"
        notFoundContent, err := os.ReadFile("./public/guides/not-found.html")
        if err != nil {
            // TODO: Redirect to the home page maybe
        }
        partContent = notFoundContent
	}

	chaptersInfo := make(map[string][]partData)
	for _, part := range guideConf.ChapterParts {
		info := partData{
			Title: part.Title,
			Path:  fmt.Sprintf("/guides/%s/%s", guideSlug, part.Slug),
		}
		chaptersInfo[part.Chapter] = append(chaptersInfo[part.Chapter], info)
	}

	return c.Render(http.StatusOK, "guide", guideData{
		Title:    guideTitle,
		Chapters: chaptersInfo,
		Content:  template.HTML(partContent),
	})
}

func Login(c echo.Context) error {
	type data struct {
		SessionData interface{}
		CSRF        string
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	// Check if user is authenticated
	if auth, ok := sess.Values["authenticated"].(bool); ok && auth {
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "login", data{
		SessionData: sess,
		CSRF:        c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
	})
}

type FormErrors map[string]string

type LoginForm struct {
	Email                string
	Password             string
	IncorrectCredentials bool

	Errors FormErrors
}

func (f *LoginForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Email == "" {
		f.Errors["Email"] = "Please enter an email."
	} else if f.IncorrectCredentials {
		f.Errors["Email"] = "Email or password is incorrect."
	}

	if f.Password == "" {
		f.Errors["Password"] = "Please enter a password."
	}

	return len(f.Errors) == 0
}

func LoginSubmit(c echo.Context) error {
	form := LoginForm{
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		IncorrectCredentials: false,
	}

	if !form.Validate() {
		return c.Redirect(http.StatusFound, "/login")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.Values["authenticated"] = true
	sess.AddFlash("You have been logged in successfully.", "flash")

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/")
}

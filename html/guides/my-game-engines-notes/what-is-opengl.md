Is a specification developed and maintained by the [Khronos Group](https://www.khronos.org/).
OpenGL specifies how should the functions be called and what should be the output. But OpenGL does not implement those functions. Is up to developers to create implementations that comply with this
specification.

The developers that are implementing the specification are usually the graphics cards manufacturers. This means that a graphics card can support the versions of OpenGL that are coded by the developers and when there is a bug on one of the functions, is usually the fault of developers that implement the specification.

### Core-profile vs Immediate mode

Immediate mode was the old way of OpenGL. It was easy to use but it abstracted many of the inner
workings and calculations. Nowadays, the immediate mode is deprecated from version 3.2 due to it being extremely inefficient.

The new way of OpenGL is called Core-profile and it is very powerful. It gives back to the developers all the control but this also means that it is more difficult to use and learn.

### State machine

OpenGL is a large state machine. A collection of variables that define how OpenGL should currently operate. The state of OpenGL is commonly referred to as the OpenGL `context`.

How it works is that we would first set some options manipulating some buffers and the render the current context. If we want to draw lines instead of triangles, we first configure OpenGL context variables to draw lines and then we tell OpenGL to draw.

### Creating a window

Before we can actually start using OpenGL, we need to create a context and an application window where we are going to draw. These operations are not the resposibility of the OpenGL library and this is why we need to take care of creating a window, defining a context, and handling user input.

There are some other libraries that we could use that take care of exactly that, for example GLUT, SDL, SFML, and GLFW. You can learn how to get started with all of those in [this article](/guides/my-c-notes/how-to-link-popular-c-libraries).

In this article we will be using [GLFW](https://www.glfw.org/) so make sure to have it installed in your computer and setup in your project.

Once both `GLAD` and `GLFW` ready to be used, we can create create a window:

```c
#include <glad/glad.h>
#include <GLFW/glfw3.h>

int main()
{
    glfwInit();
    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 3);
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);

    GLFWWindow *window = glfwCreateWindow(800, 600, "LearnOpenGL", NULL, NULL);
    if (window == NULL)
    {
        std<<cout << "Failed to create GLFW window" << std::endl;
        glfwTerminate();
        return -1;
    }
    glfwMakeContextCurrent(window);

    if (!gladLoadGLLoader((GLADloadproc)glfwGetProcAddress))
    {
        std::cout << "Failed to initilize GLAD" << std::endl;
        return -1;
    }

    glViewport(0, 0, 800, 600);

    while (!glfwWindowShouldClose(window))
    {
        glfwSwapBuffers(window);
        glfwPollEvents();
    }

    glfwTerminate();
    return 0;
}
```

Here we initialize `GLFW`, and configure it using the `glfwWindowHint` function.
Then we create a window with a `width` and `height` and check that the object
was properly initialized. After that, we make the context of our window the main context
on the current thread.

We need then GLAD to load the OpenGL functions. GLFW helps us here to determine the correct
function based on the OS we are compiling for.

The size of the viewport is also set here to instruct the dimensions where we want to draw.

We then need to create an infinite loop that will keep our window open until any signal from
the OS is triggered (`glfwWindowShouldClose` checks for a reason to close the window).

Inside of the loop, we poll for events like keyboard input or mouse movement using the function `glfwPollEvents`.

We also swap the buffer that contains the color values for each pixel in the GLFW's window. This buffer is used 
to render the current frame to the screen.

The technique uses a [double buffer](https://en.wikipedia.org/wiki/Multiple_buffering) which is used to prevent flickering while painting the next frame, the front buffer
displays the final image that is displayed in the window, and all the draw commands that happen during the frame write to
the back buffer. Once the drawing is done, we can safely swap the front and back buffers without flickering.

The last thing to do is to clean all the resources used by GLFW.

### Resizing The Window
When the user resizes the window we need to adjust the viewport to fit on the window.

```c
#include <glad/glad.h>
#include <GLFW/glfw3.h>

void frame_buffer_size_callback(GLFWwindow *window, int width, int height)
{
    glViewport(0, 0, width, height);
}

int main()
{
    ...

    glViewport(0, 0, 800, 600);
    glfwSetFramebufferSizeCallback(window, framebuffer_size_callback);

    ...

}
```

### Processing Input

We can process user input checking if the specific key was pressed on the window context.

```c
void process_input(GLFWwindow *window)
{
    if (glfwGetKey(window, GLFW_KEY_ESCAPE) == GLFW_PRESS)
    {
        glfwSetWindowShouldClose(window, true);
    }
}

int main()
{
    ...

    while (!glfwWindowShouldClose(window))
    {
        processInput(window);

        glfwSwapBuffers(window);
        glfwPollEvents();
    }
}
```

In this case, we check if the key pressed is the `ESC` key. If that is the case, then we
instruct the GLFW to close the window in the next loop check.

### Painting Into The Window

The last cool thing that we can do is to clear the screen passing a color to be 
painted with.

```c
int main()
{
    ...

    while (!glfwWindowShouldClose(window))
    {
        // Process input (game update)
        processInput(window);

        // Rendering commands
        glClearColor(0.2f, 0.3f, 0.3f, 1.0f);
        glClear(GL_COLOR_BUFFER_BIT);

        glfwSwapBuffers(window);
        glfwPollEvents();
    }
}
```

Now that we can paint a window we can finally dive into learning OpenGL and the 
inner workings. In the next article we will paint the famous triangle and even more!

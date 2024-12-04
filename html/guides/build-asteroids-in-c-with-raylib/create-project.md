# Creating The Game Project

Now that we have a C/C++ compiler installed and Raylib ready to be used, we
can finally start our project by creating a folder where the game will live.

Go to the directory where you would like to have your project in the terminal
and create a new folder:

```bash
mkdir asteroids
cd asteroids
```

We will first need an entry point to run our game. Generally, we would create
a `main.c` file with a main function that will be called when
running the compiled program. But we can call this file whatever we want as
long as it contains the main function.

One additional step, that I like to do is place all the code in a sub folder
called `src` to have a bit of structure and not have everything in
the root. I will also create a sub folder called `bin` where the
compiled code will be stored.

```bash
mkdir src bin
```

Let us create the entry point. I will create a file called <code>main.c</code>

```c
#include "stdio.h"

int main() {
    printf("Hello World");
    return 0;
}
```

We can now compile this program and run it by running the following in the
terminal [Check that you have a compiler](/guides/build-asteroids-in-c-with-raylib/setup-environment) available in your machine:

```bash
gcc ./src/main.c -o ./bin/Asteroids
```

If everything went well, you should have an executable file created inside the
`bin` called `Asteroids`.

Lets run it now.

```bash
./bin/Asteroids
```

This should have printed into the terminal `"Hello World"`

Printing `"Hello World"` is fine but we want to see more action and jump already
to code the game so let us create our game window using raylib.

```c
#include "stdio.h"
#include "raylib.h"

int main() {
    const int screen_width = 800;
    const int screen_height = 450;

    InitWindow(screen_width, screen_height, "raylib [core] example - basic window");

    SetTargetFPS(60);

    while (!WindowShouldClose()) {
        BeginDrawing();
        ClearBackground(RAYWHITE);

        DrawText("Congrats! You created your first window", 190, 200, 20, LIGHTGRAY);
        EndDrawing();
    }

    CloseWindow();

    return 0;
}
```

There are a couple of things happening here, but I, basically, just copy
pasted the most simple example from the [Raylib's examples](https://www.raylib.com/examples.html)
page. But no worries, I will explain step by step what is happening here:

1. First we included the `raylib.h` which contains all the constans
   and function signatures of raylib. It is not the actual code that runs, but
   it is a contract informing what are the functions, constants available and
   how should we use them.
2. Then we initialized the window passing a width and height. This prepares the
   window that will be created and also the [OpenGL](https://www.opengl.org/) context that
   will be used to paint the pixels.
3. We set then the target [FPS](https://en.wikipedia.org/wiki/Frame_rate) (Frames Per Second) to 60.
4. We create the game loop by using a while loop that will only stop executing
   until the signal `WindowShouldClose` is triggered. This could
   happen by just closing the window or maybe the Operating System decides to
   close the program for some reason.
5. While running, we start drawing. First we clear the background with a
   `RAYWHITE` color, and then we draw some text in the middle of the
   screen.
6. Lastly, before returning from the main function, the window has to be closed
   properly, so any reasources that where allocated can be properly released.

Seems like a lot to diggest but don't worry, this is all we need to start
coding the game finally. Raylib makes it super simple, that way we don't even
have to go deeper to try to understand how the actual pixels are being
painted.

Now is time to run our code. But first we need to compile it.

Let us try the same command that we used previously:

```bash
gcc ./src/main.c -o ./bin/Asteroids
```

Something is not working correctly. The compiler is not able to find
`raylib.h`

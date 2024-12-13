The player ship in the traditional [Asteroids](<https://en.wikipedia.org/wiki/Asteroids_(video_game)>) game is just a triangular shaped object. It looks a bit like the cursor on a computer.

[TODO]: Image of the player ship

With raylib we can construct something this by putting together a couple of lines that are connected by points (or vertex). In my case, my ship will be constructed with 6 points and I will draw lines connecting them together.

```c
#define PLAYER_POINTS = 6;

Vector2 points[PLAYER_POINTS] = {
    (Vector2){0.0f, -20.0f}, (Vector2){-20.0f, 20.0f}, (Vector2){-8.0f, 12.0f},
    (Vector2){8.0f, 12.0f},  (Vector2){20.0f, 20.0f},  (Vector2){0.0f, -20.0f},
};
```

Here we are using a struct type from raylib called `Vector2` which allows us to define
a 2D coordinate point with `x` and `y` values.

### Drawing The Lines

Now we can loop over this points array and draw lines between them with a function
provided by raylib called `DrawLineEx` (check [cheatsheet](https://www.raylib.com/cheatsheet/cheatsheet.html)) which accepts the start position (`Vector2`), the end position (`Vector2`), the thickness of the line (`float`), and the color of the line.

Let us see how this looks:

```c
for (int i = 0; i < PLAYER_POINTS; i++)
{
    DrawLineEx(points[i], points[(i + 1) % PLAYER_POINTS], 2.0f, BLACK);
}
```

This code will loop over all the points drawing lines between the current point and the next point.
When we reach to the last point we want to connect the last to the first point to close the shape, that is why we use the module operator `(i + 1) % PLAYER_POINTS` to wrap the index back to `0`.

Now, lets put everything together in our program to see our player on the screen.

```c
#include "raylib.h"

#define PLAYER_POINTS 6

int main() {
    const int screen_width = 800;
    const int screen_height = 450;

    InitWindow(screen_width, screen_height, "raylib [core] example - basic window");

    SetTargetFPS(60);

    Vector2 points[PLAYER_POINTS] = {
        (Vector2){0.0f, -20.0f}, (Vector2){-20.0f, 20.0f}, (Vector2){-8.0f, 12.0f},
        (Vector2){8.0f, 12.0f},  (Vector2){20.0f, 20.0f},  (Vector2){0.0f, -20.0f},
    };

    while (!WindowShouldClose()) {
        BeginDrawing();
            ClearBackground(RAYWHITE);

            for (int i = 0; i < PLAYER_POINTS; i++)
            {
                DrawLineEx(points[i], points[(i + 1) % PLAYER_POINTS], 2.0f, BLACK);
            }
        EndDrawing();
    }

    CloseWindow();

    return 0;
}
```

Compile your code, then run the program.

You should see part of your ship on the top left corner now like this:

[TODO]: Image of the player ship

### Centering The Ship

The ship was drawn in the top left corner because generally the origin of coordinates in
a computer starts there and then moves in the positive axis to the right and positive to the bottom.Like displayed in this picture:

[TODO]: Image of the the screen coordinate plane

Now, lets create a point in the screen where our ship will start (we can place it on the center of the screen), and translate the ship to that point in the 2D coordinates.

```c
#include "raymath.h"

...

Vector2 player_position = {.x = screen_width / 2.0f, .y = screen_height / 2.0f};

for (int i = 0; i < PLAYER_POINTS; i++) {
    Vector2 start_pos = Vector2Add(points[i], player_position);
    Vector2 end_pos = Vector2Add(points[(i + 1) % PLAYER_POINTS], player_position);

    DrawLineEx(start_pos, end_pos, 2.0f, BLACK);
}
```

Raylib has some convenient utility functions to deal with vector arithmetic.
For this, you need to create a header file called `raymath.h` and copy the contents
from [github](https://github.com/raysan5/raylib/blob/master/src/raymath.h).

Now we can add two vectors, the `player_position` and each of the points of the ship.
This will translate every point to where the player will be positioned.

If we compile and run the game we can see that the ship has moved to the center of the screen.

### Player Struct

Now let us create a player struct that will contain the player position and after we will add
some more info like the player rotation.

```c
typedef struct player {
  Vector2 position;
} player_t;

...

player_t player = {
    .position = (Vector2){
        .x = screen_width / 2.0f,
        .y = screen_height / 2.0f,
    },
};
```

The `player_t` struct will contain all the information that we need for our player.

### Player Module

I like to separate my code early to some clear modules. In this case I want to move all the code related to the player, to its own header file `player.h` and a c file `player.c`. That way we know that anything related to the player will live there and we don't end up will a huge main.c file. We will follow this pattern throughout the guide.

In general, for small games I like to have all the code in a single `game.c` and `game.h` file, but for this project we will keep it separate.

Let us create the `player.h` header file:

```c
#ifndef PLAYER_H_
#define PLAYER_H_

#include "raylib.h"

#define PLAYER_POINTS 6

typedef struct player {
  Vector2 position;
} player_t;

void player_draw(player_t *player);

#endif
```

Now lets implement the `player_draw` function in the new `player.c` file:

```c
#include "raylib.h"
#include "player.h"
#include "raymath.h"

void player_draw(player_t *player) {
    for (int i = 0; i < PLAYER_POINTS; i++) {
        Vector2 start_pos = Vector2Add(points[i], player->position);
        Vector2 end_pos = Vector2Add(points[(i + 1) % PLAYER_POINTS], player->position);

        DrawLineEx(start_pos, end_pos, 2.0f, BLACK);
    }
}
```

Now we can clean up the `main.c` file and use the player module to draw the player ship.

```c
#include "raylib.h"
#include "player.h"

int main() {
    const int screen_width = 800;
    const int screen_height = 450;

    InitWindow(screen_width, screen_height, "raylib [core] example - basic window");

    SetTargetFPS(60);

    player_t player = {
        .position = (Vector2){
            .x = screen_width / 2.0f,
            .y = screen_height / 2.0f,
        },
    };

    while (!WindowShouldClose()) {
        BeginDrawing();
            ClearBackground(RAYWHITE);

            player_draw(&player);
        EndDrawing();
    }

    CloseWindow();

    return 0;
}
```

This looks much cleaner. The only thing that we need to do in the main function is initialize the player and pass a reference to the `player_draw` function. We do this to keep the player struct in the [stack](/guides/my-c-notes/stack-vs-heap) stored until the program closes.

Now that we have a player module and we can draw the player ship, is time to let the ship move freely around the screen. We will learn in the next article how to check for user input and move the player and rotate the ship to move forward.

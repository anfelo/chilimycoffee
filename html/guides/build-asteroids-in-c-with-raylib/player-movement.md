# Player Movement

Now that we have a space ship, we can now give it some life and allow the player
to move the ship around the screen.

First let us describe what to expect about the player movement.
When the player goes forward (pressing `KEY_UP`) the ship should move in the direction
the ship is currently facing. To change direction, the player can press `KEY_LEFT` or `KEY_RIGHT`
to steer the ship to another direction. This will rotate the ship in place and the player
will be able to move forward with the pressing `KEY_UP`.

### Naive Movement

We will start very simple with the movement and slowly build up to a nice space ship like
movement. To start we first can create a `player_update` function that we can call to handle
all the user input and state changes to the player before it is painted to the screen.

Add a new function declaration and a constant `PLAYER_SPEED` in the `player.h` file:

```c
#ifndef PLAYER_H_
#define PLAYER_H_

#include "raylib.h"

#define PLAYER_POINTS 6
#define PLAYER_SPEED 2.0f

typedef struct player {
  Vector2 position;
} player_t;

void player_draw(player_t *player);
void player_update(player_t *player);

#endif
```

Now let us implement this function in the `player.c`:

```c
#include "raylib.h"
#include "player.h"
#include "raymath.h"

...

void player_update(player_t *player) {
    Vector2 velocity = {0};
    if (IsKeyDown(KEY_UP)) {
        // INFO: The ship is currently lookin up. When the arrow up key is pressed
        // it should move upwards which means we need to move with a negative velocity
        // in the y axis.
        velocity.y = -PLAYER_SPEED;
        player.position = Vector2Add(player.position, velocity);
    }
}
```

This is a bit borring, the ship can only move upwards and can't go back. For that we need
to introduce a player rotation, so we can rotate the ship with the `KEY_LEFT` and `KEY_RIGHT`.
We need to add the player rotation to the `player_t` struct to keep track of it.

```c
...

typedef struct player {
  Vector2 position;
  float rotation;
} player_t;

...
```

Now, lets update the player rotation when the arrow keys are pressed. Make sure to also
update the `player_draw` function so the ship is also drawn with the new rotation:

```c
// INFO: Make sure to update the `player.h` with this new constant
#define PLAYER_ROT_SPEED 2.0f

...

void player_draw(player_t *player) {
    // INFO: Rotate each of the points
    Vector2 start_point_rot = Vector2Rotate(points[i], player.rotation * DEG2RAD);
    Vector2 end_point_rot = Vector2Rotate(points[(i + 1) % PLAYER_POINTS],
                                        player.rotation * DEG2RAD);

    // INFO: Use the rotated points to scale them and add the player position
    // in the screen
    Vector2 start_pos = Vector2Add(Vector2Scale(start_point_rot, PLAYER_SIZE),
                                 player.position);
    Vector2 end_pos =
      Vector2Add(Vector2Scale(end_point_rot, PLAYER_SIZE), player.position);

    DrawLineEx(start_pos, end_pos, 2.0f, BLACK);
}

void player_update(player_t *player) {
    Vector2 velocity = {0};
    if (IsKeyDown(KEY_UP)) {
        velocity.y = -PLAYER_SPEED;
        player.position = Vector2Add(player.position, velocity);
    }

    float rotation = (IsKeyDown(KEY_RIGHT) - IsKeyDown(KEY_LEFT)) * PLAYER_ROT_SPEED;
    player.rotation += rotation;
}
```
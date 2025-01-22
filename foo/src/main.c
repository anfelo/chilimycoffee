#include "raylib.h"
#include "raymath.h"

#define SCREEN_WIDTH 800
#define SCREEN_HEIGHT 450

#define PLAYER_POINTS 6
#define PLAYER_SIZE 20.0f
#define PLAYER_SPEED 250
#define PLAYER_ACCELERATION 750
#define PLAYER_DRAG 0.95f
#define PLAYER_ROT_SPEED 360

#define FIELD_MIN_X (-PLAYER_SIZE)
#define FIELD_MAX_X (SCREEN_WIDTH + PLAYER_SIZE)
#define FIELD_MIN_Y (-PLAYER_SIZE)
#define FIELD_MAX_Y (SCREEN_HEIGHT + PLAYER_SIZE)

typedef struct player {
  Vector2 position;
  Vector2 velocity;
  float rotation;
} player_t;

void player_wrap(player_t *player);

int main() {
  const int screen_width = 800;
  const int screen_height = 450;

  InitWindow(screen_width, screen_height,
             "raylib [core] example - basic window");

  SetTargetFPS(60);

  Vector2 points[PLAYER_POINTS] = {
      (Vector2){0.0f, -1.0f}, (Vector2){-1.0f, 1.0f}, (Vector2){-0.4f, 0.6f},
      (Vector2){0.4f, 0.6f},  (Vector2){1.0f, 1.0f},  (Vector2){0.0f, -1.0f},
  };
  player_t player = {
      .position =
          (Vector2){.x = screen_width / 2.0f, .y = screen_height / 2.0f},
      .velocity = {0},
      .rotation = 0,
  };

  while (!WindowShouldClose()) {
    float frametime = GetFrameTime();
    Vector2 player_facing_direction =
        Vector2Rotate((Vector2){0, -1}, player.rotation * DEG2RAD);
    float magSqr = Vector2LengthSqr(player.velocity);
    float mag = sqrt(magSqr);

    if (IsKeyDown(KEY_UP)) {
      player.velocity = Vector2Add(
          player.velocity, Vector2Scale(player_facing_direction,
                                        PLAYER_ACCELERATION * frametime));
      if (mag > PLAYER_SPEED) {
        player.velocity = Vector2Scale(player.velocity, PLAYER_SPEED / mag);
      }

      player_wrap(&player);
    } else {
      if (mag > 0) {
        player.velocity = Vector2Scale(player.velocity, PLAYER_DRAG);
      }
    }

    player.position =
        Vector2Add(player.position, Vector2Scale(player.velocity, frametime));

    float x_input = (IsKeyDown(KEY_RIGHT) - IsKeyDown(KEY_LEFT));
    player.rotation += x_input * PLAYER_ROT_SPEED * frametime;

    BeginDrawing();
    ClearBackground(BLACK);

    for (int i = 0; i < PLAYER_POINTS; i++) {
      Vector2 start_point_rot =
          Vector2Rotate(points[i], player.rotation * DEG2RAD);
      Vector2 end_point_rot = Vector2Rotate(points[(i + 1) % PLAYER_POINTS],
                                            player.rotation * DEG2RAD);

      Vector2 start_pos = Vector2Add(Vector2Scale(start_point_rot, PLAYER_SIZE),
                                     player.position);
      Vector2 end_pos =
          Vector2Add(Vector2Scale(end_point_rot, PLAYER_SIZE), player.position);

      DrawLineEx(start_pos, end_pos, 2.0f, RAYWHITE);
    }

    EndDrawing();
  }

  CloseWindow();

  return 0;
}

void player_wrap(player_t *player) {
  if (player->position.x > FIELD_MAX_X) {
    player->position.x = -PLAYER_SIZE;
  } else if (player->position.x < FIELD_MIN_X) {
    player->position.x = SCREEN_WIDTH + PLAYER_SIZE;
  }

  if (player->position.y > FIELD_MAX_Y) {
    player->position.y = -PLAYER_SIZE;
  } else if (player->position.y < FIELD_MIN_Y) {
    player->position.y = SCREEN_HEIGHT + PLAYER_SIZE;
  }
}

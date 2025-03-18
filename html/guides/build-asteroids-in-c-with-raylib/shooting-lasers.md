Now that our ship moves freely around the window, we can add a bit of action to the
scene. First let us add to capability to shoot intergalactic lasers to the ship.
We don't have yet what to shoot at but that will come shortly.

Let us start first by creating a projectile struct that will contain all
the information that we need to spawn projectiles.

```c
typedef struct projectile {
  Vector2 position;
  float rotation;
  float creationTime;
  bool active;
} projectile_t;
```

Now, the player will be shooting many of these projectiles, some will hit asteroids but some others will fly out of the window and be destroyed after some time.

We can do something smart here by creating an object pool for projectiles and activating them when the player shoots and deactivating them when they hit an asteroid or they go
off screen.

```c
#define PROJECTILE_MAX 12;

static projectile_t _projectiles[PROJECTILE_MAX];
```

Here we create a fixed length array of 12 projectiles that we can reuse.

Now let us add handy function to add new projectiles. This function will be called
when the player fires.

```c
void add_projectile(Vector2 position, float rotation) {
  for (int i = 0; i < PROJECTILE_MAX; i++) {
    if (_projectiles[i].active) {
      continue;
    }

    _projectiles[i] = (projectile_t){
        .position = position,
        .rotation = rotation,
        .creationTime = GetTime(),
        .active = true,
    };

    return;
  }

  // Failed to create a projectile because there was no inactive spots in the array!
}
```

This function loops over all the projectiles in the object pool and creates a new 
projectile in the given position and with a given rotation. We also set it to active.

When all the projectiles are already active, we can not create more projectiles. This 
will rarely happen because we will give the player a delay between shootings so
they can't create a big amount of projectiles pressing the fire button as fast as they 
can.

```c
#define PLAYER_PROJECTILE_OFFSET PLAYER_SIZE
#define PLAYER_FIRE_DELAY 0.33f

...

void player_update(player_t *player, float frametime) {
  // player movement code
  ...

  float time = GetTime();
  if (IsKeyDown(KEY_SPACE)) {
    if (time > player->lastFireTime + PLAYER_FIRE_DELAY) {
      add_projectile(
          Vector2Add(player->position, Vector2Scale(player_facing_direction,
                                                    PLAYER_PROJECTILE_OFFSET)),
          player->rotation);
      player->lastFireTime = GetTime();
    }
  }
}
```

In the `player_update` function we check for the `KEY_SPACE` to be down to fire a 
new projectile. We also check that the current time is greater than the `lastFireTime`
plus a delay. This will the amount of times that the `add_projectile` function is
triggered.

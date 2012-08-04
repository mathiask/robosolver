#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "robosolver.h"

// The size of the field.
unsigned N = 0;

// Move from a location into a given direction.
// The field is updated in place.
// Returns the target location, -1 if the move is illegal
int doMove(field f, location from, direction d) {
  unsigned x = X(from);
  unsigned y = Y(from);
  color c = COLOR(f[from]);
  int dx = 0;
  int dy = 0;
  unsigned x2, y2;
  switch (d) {
  case NORTH: dy = -1; break;
  case SOUTH: dy = 1; break;
  case WEST: dx = -1; break;
  case EAST: dx = 1; break;
  }
  x2 = x + dx;
  y2 = y + dy;
  while (!(f[xy(x, y)] & d) && !COLOR(f[xy(x2, y2)])) {
    x = x2;
    y = y2;
    x2 += dx;
    y2 += dy;
  }
  if (x == X(from) && y == Y(from)) {
    return -1;
  }
  f[from] = WALLS(f[from]);
  f[xy(x, y)] |= c << 4;
  return xy(x, y);
}

field cloneField(field f) {
  field clone = malloc(N*N+1);
  memcpy(clone, f, N*N);
  clone[N*N] = 0;
  return clone;
}

// Recursively from a location to another location with max moves.
// Returns the path as NULL-terminated string of directions.
int moveTo(field f, location from, location to, unsigned max, direction *moves) {
  if (from == to) {
    *moves = 0;
    return 1;
  }
  if (max < 1) {
    return 0;
  }
  for (int i = 0; i < 4; ++i) {

    field f2 = cloneField(f);
    int newFrom = doMove(f2, from, 1 << i);
    *moves= 1<<i;
    if (newFrom >= 0 && moveTo(f2, newFrom, to, max - 1, moves + 1)) {
      free(f2);
      return 1;
    }
    free(f2);
  }
  return 0;
}

direction oppositeDirection(direction d) {
  switch (d) {
  case NORTH: return SOUTH;
  case SOUTH: return NORTH;
  case WEST: return EAST;
  case EAST: return WEST;
  }
}

int solve(field f, location *robot, location to, unsigned max, move *moves) {
  //  printf("max: %d, to: %d, ls: %d,%d,%d,%d\n%s", max, (int)to, (int)robot[0], (int)robot[1], (int)robot[2], (int)robot[3], prettyPrint(f));
  if (robot[0] == to) {
    (moves+1)->d = 0;
    return 1;
  }
  if (max < 1) {
    return 0;
  }
  for (int i = 0; i < 4; ++i) {
    for (int d = 0; d < 4; ++d) {
      if (moves->c == i + 1 &&  moves->d == oppositeDirection(1 << d))
        continue;
      field f2 = cloneField(f);
      location from = robot[i];
      int target = doMove(f2, from, 1 << d);
      if (target >= 0) {
        (moves+1)->c = i + 1;
        (moves+1)->d = 1<<d;
        robot[i] = target;
        if (solve(f2, robot, to, max - 1, moves + 1)) {
          free(f2);
          return 1;
        }
      }
      free(f2);
      robot[i] = from;
    }
  }
  return 0;
}

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
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

void moveCommand(field f, location from, direction d) {
  printf("%s\n", doMove(f, from, d) >= 0 ? prettyPrint(f) : "Illegal move!\n");
}

void toCommand(field f, location from, location to) {
  direction moves[100];
  if (moveTo(f, from, to, 10, moves)) {
    for (direction *d = moves; *d; d++) {
      printf("%d ", *d);
    }
    printf("\n");
  } else {
    printf("No solution!\n");
  }
}

void findCommand(field f, color c) {
  int loc = findColor(f, c);
  if (loc < 0) {
    printf("Not Found.\n");
  } else {
    printf("Color %d is at %d\n", c, loc);
  }
}

// robosolver <size> <pos> move <from> <direction>
// robosolver <size> <pos> to <from> <to>
// robosolver <size> <pos> find <color>
int main(int argc, const char** argv) {
    N = atoi(argv[1]);
    field f = parse(strdup(argv[2]));
    if (strcmp("move", argv[3]) == 0) {
      moveCommand(f, atoi(argv[4]), atoi(argv[5]));
    } else if (strcmp("to", argv[3]) == 0) {
      toCommand(f, atoi(argv[4]), atoi(argv[5]));
    } else if (strcmp("find", argv[3]) == 0) {
      findCommand(f, atoi(argv[4]));
    }

    return 0;
}

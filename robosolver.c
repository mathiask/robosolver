#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "robosolver.h"

// The size of the field.
unsigned N = 0;

// Move from a location into a given direction.
// The field is updated in place.
// Returns the field after the move.
// NULL if the move is illegal
field doMove(field f, location from, direction d) {
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
    return NULL;
  }
  f[from] = WALLS(f[from]);
  f[xy(x, y)] |= c << 4;
  return f;
}

// Recursively from a location to another location with max moves.
// Returns the path as NULL-terminated string of directions.
direction* moveTo(field f, location from, location to, unsigned max);

// robosolver <size> <pos> move <from> <direction>
int main(int argc, const char** argv) {
    N = atoi(argv[1]);
    const char *pos = strdup(argv[2]);
    location from = atoi(argv[4]);
    direction d = atoi(argv[5]);
    field f = parse(pos);
    printf("%s\n", prettyPrint(f));
    field resultPos = doMove(f, from, d);
    printf("%s\n", resultPos ? prettyPrint(resultPos) : "Illegal move!\n");
    return 0;
}

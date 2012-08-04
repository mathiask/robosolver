#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "robosolver.h"

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

char *prettyDirection(direction d) {
  switch (d) {
  case NORTH: return "north";
  case SOUTH: return "south";
  case WEST: return "west";
  case EAST: return "east";
  }
}

void solveCommand(field f, location to) {
  location robot[4];
  move moves[100];
  moves->c = 0;
  moves->d = 0;
  for (int i = 1; i <= 4; i++) {
    robot[i - 1] = findColor(f, i);
  }
  for (int depth = 1; depth < 20; depth++) {
    if (solve(f, robot, to, depth, moves)) {
      for (move *m = moves + 1; m->d; m++) {
        printf("%d: %s\n", m->c, prettyDirection(m->d));
      }
      break;
    } else {
      printf("No solution at depth %d!\n", depth);
    }
  }
}

// robosolver <size> <pos> move <from> <direction>
// robosolver <size> <pos> to <from> <to>
// robosolver <size> <pos> find <color>
// robosolver <size> <pos> solve <to>
int main(int argc, const char** argv) {
    N = atoi(argv[1]);
    field f = parse(strdup(argv[2]));
    if (strcmp("move", argv[3]) == 0) {
      moveCommand(f, atoi(argv[4]), atoi(argv[5]));
    } else if (strcmp("to", argv[3]) == 0) {
      toCommand(f, atoi(argv[4]), atoi(argv[5]));
    } else if (strcmp("find", argv[3]) == 0) {
      findCommand(f, atoi(argv[4]));
    } else if (strcmp("solve", argv[3]) == 0) {
      solveCommand(f, atoi(argv[4]));
    }

    return 0;
}

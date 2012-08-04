#include <stdio.h>
#include <string.h>
#include "robosolver.h"

// The size of the field.
unsigned N = 0;

// Move from a location into a given direction.
// Returns the field after the move.
unsigned char* doMove(field f, location from, direction d);

// Recursively from a location to another location with max moves.
// Returns the path as NULL-terminated string of directions.
unsigned char* moveTo(field f, location from, location to, unsigned max);

// robosolver <size> <pos> <color> <destination>
int main(int argc, const char** argv) {
    N = atoi(argv[1]);
    field pos = strdup(argv[2]);
    color col = atoi(argv[3]);
    location dest = atoi(argv[4]);
    return 0;
}

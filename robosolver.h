// Walls
#define NORTH    (1<<0)
#define EAST     (1<<1)
#define SOUTH    (1<<2)
#define WEST     (1<<3)

// Colors
#define EMPTY    0
#define BLUE     1
#define GREEN    2
#define YELLOW   3
#define RED      4

#define COLOR(x) (((x) >> 4) & 0x07)
#define WALLS(x) ((x) & 0x0f)

#define X(location, n) ((location) % (n))
#define Y(location, n) ((lcoation) / (n))

typedef unsigned char point;
typedef unsigned char direction;
typedef unsigned char* field;
typedef unsigned location;
typedef unsigned char color;

// Move from a location into a given direction.
// Returns the field after the move.
unsigned char* doMove(field f, location from, direction d);

// Recursively from a location to another location with max moves.
// Returns the path as NULL-terminated string of directions.
unsigned char* moveTo(field f, location from, location to, unsigned max);

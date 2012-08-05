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

// The size of the field.
extern unsigned N;

#define X(location) ((location) % (N))
#define Y(location) ((location) / (N))
#define xy(X, Y) ((Y)*N+(X))

typedef unsigned char point;
typedef unsigned char direction;
typedef unsigned char* field;
typedef unsigned location;
typedef unsigned char color;

typedef struct move {
  color c;
  direction d;
} move;

typedef struct hashEntry {
    struct hashEntry* next;
    unsigned char robots[4];
    unsigned remainingDepth;
} hashEntry;

int lookup(location* robot, unsigned remainingDepth);

// Move from a location into a given direction.
// The field is updated in place.
// Returns the target location, -1 if the move is illegal
int doMove(field f, location from, direction d);

// Recursively from a location to another location with max moves.
// Returns non-zero if a solution was found
int moveTo(field f, location from, location to, unsigned max, direction *moves);

// Read internal field representation from hex string.
field parse(const char *positionString);

char *prettyPrint(field f);

// -1 when not found
int findColor(field f, color c);

// Solve the problem of moving robot 1 to target
int solve(field f, location *robot, location to, unsigned max, move *moves);

package core

// Walls and directions.
type Direction byte
const (
	NORTH Direction = 1 << iota
	EAST
	SOUTH
	WEST
)

// One byte encoding the walls around the square in the bottom four bits
// and the robot on the square in the top four bits.
// (Robosolver's unused "point" type.)
type Square byte

// "func (s Square) Color() byte" is also possible
func Color(s Square) byte {
	return byte((s >> 4) & 7)
}

func EncodeColor(c byte) Square {
	return Square(c << 4)
}

func Walls(s Square) byte {
	return byte(s & 0x0f)
}

// Let's hope the compiler optimizes this as NOP :-)
func Wall(d Direction) Square {
	return Square(d)
}

// A board of size n, containing a slice of n^2 Squares.
type Board struct {
	size uint
	field []Square
}

func NewBoard(n uint) *Board {
	return &Board{n, make([]Square, n * n)}
}

// Offset into the field slice.
type Location uint

// This is NOT Robosolver's xy(), but the pair (x(), y()).
func (b *Board) XY(l Location) (uint, uint) {
	return uint(l) % b.size, uint(l) / b.size
}

// Robosolvers's xy()
func (b *Board) Location(x, y uint) Location {
	return Location(y * b.size + x)
}

func (b *Board) fieldAt(x, y uint) *Square {
	return &b.field[b.Location(x, y)]
}

func (b *Board) Color(l Location) byte {
	return Color(b.field[l])
}

// Move from a location into a given direction.
// The board is updated in place and returns the target location if the
// robot can make at least one step.
// (Robosolver's doMove).
func (b *Board) moveToWall(from Location, direction Direction) (Location, bool) {
	delta := b.delta(direction)
	for at := int(from) + delta; false; at += delta {}

	return 0, true
}

func (b *Board) delta(d Direction) int {
	switch (d) {
	case NORTH: return - int(b.size)
	case SOUTH: return int(b.size)
	case WEST: return -1
	case EAST: return 1
	}
	return 0;
}

package core

// One byte encoding the walls around the square in the bottom four bits
// and the robot on the square in the top four bits.
type Square byte

// "func (s Square) Color() byte" is also possible
func Color(s Square) byte {
	return byte((s >> 4) & 7)
}

func Walls(s Square) byte {
	return byte(s & 0x0f)
}

// A board of size n containing a slice of n^2 Squares.
type Board struct {
	size uint
	field []Square
}

func NewBoard(n uint) *Board {
	return &Board{n, make([]Square, n * n)}
}

// Offset into the field slice.
type Location uint

func (b *Board) X(l Location) uint {
	return uint(l) % b.size
}

func (b *Board) Y(l Location) uint {
	return uint(l) / b.size
}

func (b *Board) XY(x, y uint) Location {
	return Location(y * b.size + x)
}

// Walls and directions.
type Direction byte
const (
	NORTH Direction = 1 << iota
	EAST
	SOUTH
	WEST
)

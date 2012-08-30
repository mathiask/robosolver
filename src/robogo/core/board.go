package core

// One byte encoding the walls around the square in the bottom four bits
// and the robot on the square in the top four bits.
type Square byte

// A board of size n containing a slice of n^2 Squares.
type Board struct {
	size uint
	field []Square
}

func NewBoard(n uint) *Board {
	return &Board{n, make([]Square, n * n)}
}

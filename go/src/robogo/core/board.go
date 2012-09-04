package core

// Walls and directions.
type Direction byte
const (
	NORTH Direction = 1 << iota
	EAST
	SOUTH
	WEST
)

func OppositeDirection(d Direction) Direction {
	switch (d) {
	case NORTH: return SOUTH
	case SOUTH: return NORTH
	case WEST: return EAST
	case EAST: return WEST
	}
	panic(d);
}

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

func Walls(s Square) Square {
	return s & 0x0f
}

// Let's hope the compiler optimizes this as NOP :-)
func Wall(d Direction) Square {
	return Square(d)
}

// A board of size n, containing a slice of n^2 Squares.
type Board struct {
	size uint
	field []Square
	undoBuffer []step
}

type step struct { from, to Location }

func (s *step) reverse() *step {
	s.from, s.to = s.to, s.from
	return s
}

func NewBoard(n uint) *Board {
	return &Board{n, make([]Square, n * n), nil}
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

// Empty board with walls around the otside.
func NewWalledBoard(n uint) *Board {
	b := Board{n, make([]Square, n * n), nil}
	for i := uint(0); i < n; i++ {
		*b.fieldAt(i, 0)     |= Wall(NORTH)
		*b.fieldAt(i, n - 1) |= Wall(SOUTH)
		*b.fieldAt(0, i)     |= Wall(WEST)
		*b.fieldAt(n - 1, i) |= Wall(EAST)
	}
	return &b
}

func (b *Board) Color(l Location) byte {
	return Color(b.field[l])
}

// Move from a location into a given direction.
// The board is updated in place and returns the target location if the
// robot can make at least one step.
// (Robosolver's doMove).
func (b *Board) MoveToWall(from Location, direction Direction) (Location, bool) {
	to := b.findWall(from, direction)
	if to == from {
		return 0, false
	}
	step := step{from, to}
	b.move(&step)
	b.undoBuffer = append(b.undoBuffer, step)
	return to, true
}

func (b *Board) findWall(from Location, direction Direction) Location {
	delta := b.delta(direction)
	to, next := from, Location(int(from) + delta)
	for b.field[to] & Wall(direction) == 0 && Color(b.field[next]) == 0 {
		to = next
		next = Location(int(next) + delta)
	}
	return to
}

func (b *Board) delta(d Direction) int {
	switch (d) {
	case NORTH: return - int(b.size)
	case SOUTH: return int(b.size)
	case WEST: return -1
	case EAST: return 1
	}
	panic(d);
}

func (b *Board) move(s *step) {
	old := b.field[s.from]
	robot := Color(old)
	b.field[s.from] = Walls(old) // remove robot
	b.field[s.to] |= EncodeColor(robot)
}

func (b *Board) Undo() {
	n := len(b.undoBuffer)
	b.move(b.undoBuffer[n - 1].reverse())
	b.undoBuffer = b.undoBuffer[:n-1]
}

// Public functions for GUI

func (b *Board) Size() uint {
	return b.size
}

func (b *Board) WallsAt(x, y uint) Square {
	return Walls(*b.fieldAt(x, y))
}

// Set robots to defined positions.
// Returns self for chaining.
func (b *Board) Reset(robots *[4][2]uint) *Board {
	for i := range b.field { b.field[i] = Walls(b.field[i]) }
	for i := byte(0); i < 4; i++ {
		*b.fieldAt(robots[i][0], robots[i][1]) |= EncodeColor(i + 1)
	}
	return b
}

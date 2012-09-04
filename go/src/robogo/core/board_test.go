package core

import (
	"testing"
	"fmt"
)

func TestBoardConstructor(t *testing.T) {
	const n, n2 = 5, 25
	b := NewBoard(n)
	if s := b.size; s != n {
		t.Errorf("Size was %v, expected %v", s, n)
	}
	if l := len(b.field); l != n2 {
		t.Errorf("Number of fields was %v, expected %v", l, n2)
	}
}

func ExampleColor() {
	const color, walls = 3, 2
	fmt.Println(Color(16 * color + walls))
	// Output:
	// 3
}

func ExampleWalls() {
	const color, walls = 3, 2
	fmt.Println(Walls(16 * color + walls))
	// Output:
	// 2
}

func ExampleBoard_XY() {
	b := NewBoard(5)
	fmt.Println(b.XY(17))
	// Output:
	// 2 3
}

func ExampleBoard_Location() {
	b := NewBoard(5)
	fmt.Println(b.Location(2, 3))
	// Output:
	// 17
}

func ExampleBoard_Color() {
	b := NewBoard(5)
	l := b.Location(2, 2)
	b.field[l] = EncodeColor(2)
	fmt.Println(b.Color(l))
	// Output:
	// 2
}

func ExampleDirection_west() {
	fmt.Println(WEST)
	// Output:
	// 8
}

func ExampleBoard_MoveToWall_onEmptyBoard() {
	b := empty7by7Board()
	start := b.Location(3, 3)
	fmt.Println(b.MoveToWall(start, NORTH)) // 0 + 3 = 3
	fmt.Println(b.MoveToWall(start, SOUTH)) // 6 * 7 + 3 = 45
	fmt.Println(b.MoveToWall(start, WEST))  // 3 * 7 + 0 = 21
	fmt.Println(b.MoveToWall(start, EAST))  // 3 * 7 + 6 = 27
	fmt.Println(b.MoveToWall(b.Location(6, 1), EAST))
	// Output:
	// 3 true
	// 45 true
	// 21 true
	// 27 true
	// 0 false
}

func empty7by7Board() *Board {
	return NewWalledBoard(7);
}

func ExampleBoard_Undo() {
	b := empty7by7Board()
	*b.fieldAt(3, 3) |= EncodeColor(1)
	b.MoveToWall(b.Location(3, 3), WEST)
	b.MoveToWall(b.Location(0, 3), NORTH)
	fmt.Println(b.field)
	b.Undo()
	b.Undo()
	fmt.Println(b.field)
	// Output:
	// [25 1 1 1 1 1 3 8 0 0 0 0 0 2 8 0 0 0 0 0 2 8 0 0 0 0 0 2 8 0 0 0 0 0 2 8 0 0 0 0 0 2 12 4 4 4 4 4 6]
	// [9 1 1 1 1 1 3 8 0 0 0 0 0 2 8 0 0 0 0 0 2 8 0 0 16 0 0 2 8 0 0 0 0 0 2 8 0 0 0 0 0 2 12 4 4 4 4 4 6]
}

func ExampleBoard_Reset() {
	b := NewBoard(3)
	*b.fieldAt(1, 1) |= EncodeColor(1)
	b.Reset(&[4][2]uint{{0, 0}, {1, 0}, {2, 0}, {3, 0}})
	fmt.Println(b.field)
	// Output:
	// [16 32 48 64 0 0 0 0 0]
}

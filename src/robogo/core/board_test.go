package core

import "testing"
import "fmt"

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

// func TestMoveToWallOnEmptyBoard(t *testing.T) {
// 	b := empty7by7Board()
// 	txy := b.Location(3, 0)
// 	if xy, ok := b.moveToWall(b.Location(3, 3), NORTH); txy != xy || !ok{
// 		t.Errorf("Moved to %v (%v), expected %v", xy, ok, txy)
// 	}
// 	txy = b.Location(3, 6)
// 	if xy, _ := b.moveToWall(b.Location(3, 3), SOUTH); txy != xy {
// 		t.Errorf("Moved to %v, expected %v", xy, txy)
// 	}
// 	txy = b.Location(0, 3)
// 	if xy, _ := b.moveToWall(b.Location(3, 3), WEST); txy != xy {
// 		t.Errorf("Moved to %v, expected %v", xy, txy)
// 	}
// 	txy = b.Location(6, 3)
// 	if xy, _ := b.moveToWall(b.Location(3, 3), EAST); txy != xy {
// 		t.Errorf("Moved to %v, expected %v", xy, txy)
// 	}
// 	if _, ok := b.moveToWall(b.Location(6, 1), EAST); ok {
// 		t.Error("Expected not OK")
// 	}
// }

func empty7by7Board() *Board {
	b := NewBoard(7)
	for i := uint(0); i < 7; i++ {
		*b.fieldAt(i, 0) |= Wall(NORTH)
		*b.fieldAt(i, 6) |= Wall(SOUTH)
		*b.fieldAt(0, i) |= Wall(WEST)
		*b.fieldAt(6, i) |= Wall(EAST)
	}
	return b
}
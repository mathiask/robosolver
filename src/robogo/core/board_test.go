package core

import "testing"

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

func TestColorOfSquare(t *testing.T) {
	const color, walls = 3, 2
	const x = 16 * color + walls
	if c := Color(x); c != color {
		t.Errorf("Color was %v, expected %v", c, color)
	}
	if w := Walls(x); w != walls {
		t.Errorf("Walls was %v, expected %v", w, walls)
	}
}

func TestX(t *testing.T) {
	b := NewBoard(5)
	x17 := uint(2)
	if x := b.X(17); x != x17 {
		t.Errorf("X was %v, expected %v", x, x17)
	}
}

func TestY(t *testing.T) {
	b := NewBoard(5)
	y17 := uint(3)
	if y := b.Y(17); y != y17 {
		t.Errorf("Y was %v, expected %v", y, y17)
	}
}

func TestXY(t *testing.T) {
	b := NewBoard(5)
	xy := Location(17)
	if n := b.XY(2, 3); xy != n {
		t.Errorf("XY was %v, expected %v", n, xy)
	}
}

func TestWEST(t *testing.T) {
	west := Direction(8)
	if WEST != west {
		t.Errorf("WEST was %v, expected %v", WEST, west)
	}
}
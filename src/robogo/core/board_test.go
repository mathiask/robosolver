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
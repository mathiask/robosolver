package core

import "fmt"

func ExampleNewPosition() {
	b := NewBoard(3)
	*b.fieldAt(0, 0) |= EncodeColor(1)
	*b.fieldAt(1, 0) |= EncodeColor(2)
	*b.fieldAt(2, 0) |= EncodeColor(3)
	*b.fieldAt(2, 2) |= EncodeColor(4)
	p := NewPosition(b, 0)
	fmt.Println(p.robot)
	// Output:
	// [0 1 2 8]
}

func ExampleSolve_target8() {
	p := NewPosition(example3x3(), 8)
	ok := p.Solve(3)
	fmt.Println(ok)
	fmt.Println(p.move[1:])
	// Output:
	// true
	// [{1 2} {1 4}]
}

func example3x3() *Board {
	b := NewWalledBoard(3)
	*b.fieldAt(0, 0) |= EncodeColor(1)
	*b.fieldAt(0, 1) |= EncodeColor(2)
	*b.fieldAt(0, 2) |= EncodeColor(3)
	*b.fieldAt(1, 2) |= EncodeColor(4)
	return b;
}

func ExampleSolve_target4() {
	p := NewPosition(example3x3(), 4)
	ok := p.Solve(3)
	fmt.Println(ok)
	fmt.Println(p.move[1:])
	// Output:
	// true
	// [{2 2} {1 4} {1 2}]
}

func ExampleSolve_target8movingPiecesOutOfTheWay() {
	b := example3x3()
	b.MoveToWall(3, EAST)
	p := NewPosition(b, 8)
	ok := p.Solve(3)
	fmt.Println(ok)
	fmt.Println(p.move[1:])
	// Output:
	// true
	// [{1 2} {2 8} {1 4}]
}

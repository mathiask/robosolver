package core

type Move struct {
	color byte
	direction Direction
}

type Position struct {
	board *Board
	robot [4]Location
	target Location
	move []Move
	hash map[uint32]hashEntry
}

type hashEntry struct {
	robot [4]Location
	remainingDepth uint
}

// This, essentially, contains Robosolver's findColor.
func NewPosition(b *Board, target Location) *Position {
	p := Position{
		b,
		*new([4]Location),
		target,
		make([]Move, 1 , 100),
		make(map[uint32]hashEntry),
	}
	p.move[0] = Move {}
	for i, x := range b.field {
		if c := Color(x); c > 0 {
			p.robot[c - 1] = Location(i)
		}
	}
	return &p
}

func (p *Position) Solve(max uint) bool {
	if p.robot[0] == p.target { return true }
	if max < 1 || lookup(p.hash, p.robot, max) { return false }
	n := len(p.move)
	lastMove := &p.move[n - 1]
	for i := byte(0); i < 4; i++ {
		for d := Direction(1); d <= (1 << 3); d <<= 1 {
			if lastMove.color == i + 1 &&
				lastMove.direction == OppositeDirection(d) {
				continue
			}
			from := p.robot[i]
			if to, ok := p.board.MoveToWall(from, d); ok {
				p.robot[i] = to
				p.move = append(p.move, Move{i + 1, d})
				if p.Solve(max - 1) {
					return true
				}
				p.robot[i] = from
				p.board.Undo()
				p.move = p.move[:n]
			}
		}
	}
	return false
}

// Lookup an make new entry if necessary.
// Returns true to indicate that the position is known to be unsolvable.
func lookup(m map[uint32]hashEntry, robot [4]Location, max uint) bool {
	return false
}
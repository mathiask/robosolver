package core

import "fmt"

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
	findRobots(&p)
	return &p
}

func findRobots(p *Position) {
	for i, x := range p.board.field {
		if c := Color(x); c > 0 {
			p.robot[c - 1] = Location(i)
		}
	}
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
	h := hash(robot)
	entry, ok := m[h]
	if !ok || entry.robot != robot {
		m[h] = hashEntry{robot, max}
		return false
	}
	if max <= entry.remainingDepth {
		return true
	}
	entry.remainingDepth = max
	return false
}

func hash(a [4]Location) uint32 {
	var result uint32
	for _, x := range a {
		result = 37 * result + uint32(x)
	}
	return result & ((1 << 24) - 1)
}

// Public functions for GUI

func (p *Position)Move() []string {
	result := make([]string, len(p.move) - 1)
	for i, m := range p.move[1:] {
		result[i] = fmt.Sprintf("%v:%s", m.color, directionName(m.direction))
	}
	return result
}

func directionName(d Direction) string {
	switch (d) {
	case NORTH: return "north"
	case SOUTH: return "south"
	case WEST: return "west"
	case EAST: return "east"
	}
	panic(d);
}

// Reset board, recompute robot locations and truncate move slice.
// Target and hash are unchanged.
func (p *Position)Reset(robots *[4][2]uint) {
	p.board.Reset(robots)
	p.move = p.move[0:1]
	findRobots(p)
}

package shared

import (
	"math/rand"
)

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func NewRandomPosition(maxN int) Position {
	return Position{rand.Intn(maxN), rand.Intn(maxN)}
}

var neighbors = []struct{ dx, dy int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (p Position) ForEachNeighbor(fn func(pos Position)) {
	for _, neighbor := range neighbors {
		neighbor := p.offset(neighbor.dx, neighbor.dy)
		fn(neighbor)
	}
}

func (p Position) offset(dx, dy int) Position {
	return Position{p.X + dx, p.Y + dy}
}

func (p Position) ForEachNeighborSatisfy(fn func(pos Position) (ok bool)) bool {
	for _, dir := range neighbors {
		neighbor := p.offset(dir.dx, dir.dy)
		if !fn(neighbor) {
			return false
		}
	}
	return true
}

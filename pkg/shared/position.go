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

var Directions = []struct{ dx, dy int }{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (p Position) ForEachDirection(fn func(pos Position)) {
	for _, dir := range Directions {
		neighbor := p.offset(dir.dx, dir.dy)
		fn(neighbor)
	}
}

func (p Position) offset(dx, dy int) Position {
	return Position{p.X + dx, p.Y + dy}
}

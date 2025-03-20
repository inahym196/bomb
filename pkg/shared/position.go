package shared

import (
	"math/rand"
	"slices"
)

type Position struct {
	X int
	Y int
}

func (pos Position) IsInside(maxX, maxY int) bool {
	return 0 <= pos.Y && pos.Y < maxY && 0 <= pos.X && pos.X < maxX
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func NewRandomPosition(maxN int) Position {
	return Position{rand.Intn(maxN), rand.Intn(maxN)}
}

func NewUniqueRandomPositionsWithout(n, maxN int, except Position) []Position {
	poss := make([]Position, n+1)
	poss[0] = except
	cnt := 1
	for cnt < n+1 {
		pos := NewRandomPosition(maxN)
		if !slices.Contains(poss, pos) {
			poss[cnt] = pos
			cnt++
		}
	}
	poss = poss[1:]
	return poss
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

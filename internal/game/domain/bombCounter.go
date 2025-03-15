package domain

import "github.com/inahym196/bomb/pkg/shared"

type bombCounter struct {
	counts [][]int
}

func (bc *bombCounter) contains(pos shared.Position) bool {
	n := len(bc.counts)
	return pos.IsInside(n, n)
}

func newBombCounter(width int) *bombCounter {
	return &bombCounter{
		counts: initCounts(width),
	}
}

func initCounts(width int) [][]int {
	counts := make([][]int, width)
	for i := range width {
		counts[i] = make([]int, width)
		for j := range width {
			counts[i][j] = 0
		}
	}
	return counts
}

func (bc *bombCounter) GetBombCounts() [][]int {
	return bc.counts
}

func (bc *bombCounter) SetBombCount(pos shared.Position) {
	bc.counts[pos.Y][pos.X] = -1
	bc.incrementBombCountForEachNeighbor(pos)
}

func (bc *bombCounter) incrementBombCountForEachNeighbor(pos shared.Position) {
	pos.ForEachNeighbor(func(p shared.Position) {
		if bc.contains(p) && bc.counts[p.Y][p.X] != -1 {
			bc.counts[p.Y][p.X]++
		}
	})
}

func (bc *bombCounter) IsNeighborsSafe(pos shared.Position) bool {
	return bc.counts[pos.Y][pos.X] == 0
}

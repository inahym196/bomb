package domain

import (
	"slices"

	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

const (
	SolutionResultIsSafe byte = iota
	SolutionResultIsBomb byte = iota
)

type Solution struct {
	Positions []shared.Position
	Result    byte
}

type Theorem interface {
	GetDescription() string
	Apply([][]interactor.CellDTO) Solution
}

type theorem1 struct{}

func (t theorem1) GetDescription() string {
	return "closedCellがBombCount以下なら全部bomb"
}

func (t theorem1) Apply(cells [][]interactor.CellDTO) Solution {
	poss := make([]shared.Position, 0, len(cells)*len(cells))
	for i := range cells {
		for j := range cells[i] {
			if !cells[i][j].IsOpened {
				continue
			}
			closedPositions := make([]shared.Position, 0, 8)
			for _, pos := range shared.NewPosition(j, i).Neighbors() {
				if !pos.IsInside(len(cells), len(cells)) || cells[pos.Y][pos.X].IsOpened {
					continue
				}
				closedPositions = append(closedPositions, pos)
			}
			if len(closedPositions) <= cells[i][j].BombCount {
				poss = append(poss, closedPositions...)
			}
		}
	}
	return Solution{Result: SolutionResultIsBomb, Positions: poss}
}

type theorem2 struct{}

func (t theorem2) GetDescription() string {
	return "周囲のflaggedCellの合計数とbombCountが等しいなら残りのclosedCellは安全"
}

func (t theorem2) Apply(cells [][]interactor.CellDTO) Solution {
	poss := make([]shared.Position, 0, len(cells)*len(cells))
	for i := range cells {
		for j := range cells[i] {
			if !cells[i][j].IsOpened {
				continue
			}
			unflaggedPositions := make([]shared.Position, 0, 8)
			flaggedCount := 0
			for _, pos := range shared.NewPosition(j, i).Neighbors() {
				if !pos.IsInside(len(cells), len(cells)) || cells[pos.Y][pos.X].IsOpened {
					continue
				}
				if cells[pos.Y][pos.X].IsFlagged {
					flaggedCount++
					continue
				}
				unflaggedPositions = append(unflaggedPositions, pos)
			}
			if flaggedCount == cells[i][j].BombCount {
				poss = append(poss, unflaggedPositions...)
			}
		}
	}
	return Solution{Result: SolutionResultIsSafe, Positions: poss}
}

type theorem3 struct{}

func (t theorem3) GetDescription() string {
	return "1・2の定理"
}

func (t theorem3) Apply(cells [][]interactor.CellDTO) Solution {

	poss := make([]shared.Position, 0, len(cells)*len(cells))
	for i := range cells {
		for j := range cells[i] {
			cell := cells[i][j]
			if !cell.IsOpened || cell.BombCount != 1 {
				continue
			}
			nbs := shared.NewPosition(j, i).Neighbors()
			for _, nb := range nbs {
				if !nb.IsInside(len(cells), len(cells)) || !cells[nb.Y][nb.X].IsOpened {
					continue
				}
				bc := cells[nb.Y][nb.X].BombCount
				if bc <= 1 {
					continue
				}
				closedDiff := make([]shared.Position, 0, 8)
				for _, diff := range t.diff(nb.Neighbors(), nbs) {
					if !cells[diff.Y][diff.X].IsOpened {
						closedDiff = append(closedDiff, diff)
					}
				}
				if len(closedDiff)+1 == bc {
					poss = append(poss, closedDiff...)
				}
			}
		}
	}
	return Solution{Result: SolutionResultIsBomb, Positions: poss}
}

func (t theorem3) diff(slice1, slice2 []shared.Position) []shared.Position {
	diff := make([]shared.Position, 0, 8)
	for _, v := range slice1 {
		if !slices.Contains(slice2, v) {
			diff = append(diff, v)
		}
	}
	return diff
}

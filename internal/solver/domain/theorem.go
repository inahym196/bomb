package domain

import (
	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

const (
	SolutionResultIsNotBomb byte = iota
	SolutionResultIsBomb    byte = iota
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
			closedPositions := make([]shared.Position, 0, 8)
			shared.NewPosition(j, i).ForEachNeighbor(func(pos shared.Position) {
				if !pos.IsInside(len(cells), len(cells)) || cells[pos.Y][pos.X].IsOpened {
					return
				}
				closedPositions = append(closedPositions, pos)
			})
			cell := cells[i][j]
			if len(closedPositions) <= cell.BombCount {
				poss = append(poss, closedPositions...)
			}
		}
	}
	return Solution{Result: SolutionResultIsBomb, Positions: poss}
}

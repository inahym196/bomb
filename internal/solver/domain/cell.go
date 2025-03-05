package domain

import (
	"github.com/inahym196/bomb/internal/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

// TODO: Impl bombScope
//type bombScope struct {
//	origin      *Cell
//	closedCells []Cell
//}
//
//func NewBombScope(cell *Cell) *bombScope {
//	return &bombScope{cell, []Cell{}}
//}

type Cell struct {
	isSolved  bool
	isClosed  bool
	bombCount int
	// TODO: bombScope *bombScope
	// TODO: scopedBy []bombScope
}

func (c Cell) IsSolved() bool    { return c.isSolved }
func (c Cell) IsClosed() bool    { return c.isClosed }
func (c Cell) GetBombCount() int { return c.bombCount }

func ToSolverCells(cells [][]interactor.CellDTO) map[shared.Position]Cell {
	scells := make(map[shared.Position]Cell)
	for i := range cells {
		for j, cell := range cells[i] {
			pos := shared.NewPosition(j, i)
			if cell.IsOpened && cell.BombCount == 0 {
				continue
			}
			scells[pos] = NewCell(cells, pos)
		}
	}
	return scells
}

func NewCell(cells [][]interactor.CellDTO, pos shared.Position) Cell {
	if !cells[pos.Y][pos.X].IsOpened {
		return Cell{false, true, -1}
	}
	return Cell{isSolvedCell(cells, pos), false, cells[pos.Y][pos.X].BombCount}
}

func isSolvedCell(cells [][]interactor.CellDTO, pos shared.Position) bool {
	// 周りのセルのうち、Cells内に含まれているセルが全てOpenの場合はSolved
	return pos.ForEachDirectionSatisfy(func(pos shared.Position) (ok bool) {
		// posがCells内に含まれない場合はSolveに関係ないためok判定
		return !isInCells(cells, pos) || cells[pos.Y][pos.X].IsOpened
	})
}

func isInCells(cells [][]interactor.CellDTO, pos shared.Position) bool {
	min := 0
	max := len(cells)
	return min <= pos.X && pos.X < max && min <= pos.Y && pos.Y < max
}

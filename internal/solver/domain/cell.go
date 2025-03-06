package domain

import (
	"maps"
	"slices"

	"github.com/inahym196/bomb/internal/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

type OpenCell struct {
	shadyCells map[shared.Position]struct{}
	bombCount  int
	//isSolved  bool
}

func (o OpenCell) GetShadyCellKeys() []shared.Position {
	return slices.Collect(maps.Keys(o.shadyCells))
}
func (o OpenCell) GetBombCount() int { return o.bombCount }

func NewSolverCells(cells [][]interactor.CellDTO) map[shared.Position]OpenCell {
	openCells := make(map[shared.Position]OpenCell)
	for i := range cells {
		for j, cell := range cells[i] {
			if cell.IsOpened && cell.BombCount != 0 {
				pos := shared.NewPosition(j, i)
				openCells[pos] = OpenCell{
					shadyCells: findShadyCellsFrom(cells, pos),
					bombCount:  cells[pos.Y][pos.X].BombCount,
				}
			}
		}
	}
	return openCells
}

func findShadyCellsFrom(cells [][]interactor.CellDTO, pos shared.Position) map[shared.Position]struct{} {
	shadyCells := make(map[shared.Position]struct{})
	pos.ForEachNeighbor(func(pos shared.Position) {
		if isInCells(cells, pos) && !cells[pos.Y][pos.X].IsOpened {
			shadyCells[pos] = struct{}{}
		}
	})
	return shadyCells
}

// そもそもSolved以外(BombCount!=0)のCellを集めているので全てのOpenCellがマッチするのでは？
// -> そのうちCell.Checkedとか実装したら活きてくる。実装してからコメントを外す
//
//	func isSolvedCell(cells [][]interactor.CellDTO, pos shared.Position) bool {
//			// 周りのセルのうち、Cells内に含まれているセルが全てOpenの場合はSolved
//			return pos.ForEachNeighborSatisfy(func(pos shared.Position) (ok bool) {
//				// posがCells内に含まれない場合はSolveに関係ないためok判定
//				return !isInCells(cells, pos) || cells[pos.Y][pos.X].IsOpened
//			})
//		}
func isInCells(cells [][]interactor.CellDTO, pos shared.Position) bool {
	min := 0
	max := len(cells)
	return min <= pos.X && pos.X < max && min <= pos.Y && pos.Y < max
}

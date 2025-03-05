package domain

import (
	"github.com/inahym196/bomb/internal/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

type ClosedCell struct {
	referencedBy []*OpenCell
}

type OpenCell struct {
	refs []*ClosedCell
	//isSolved  bool
	bombCount int
}

func ToSolverCells(cells [][]interactor.CellDTO) (map[shared.Position]OpenCell, map[shared.Position]ClosedCell) {
	openCells := make(map[shared.Position]OpenCell)
	closedCells := make(map[shared.Position]ClosedCell)
	for i := range cells {
		for j, cell := range cells[i] {
			pos := shared.NewPosition(j, i)
			if !cell.IsOpened {
				closedCells[pos] = ClosedCell{nil}
				continue
			}
			if cell.BombCount != 0 {
				openCells[pos] = NewOpenCell(cells, pos)
			}
		}
	}
	return openCells, closedCells
}

func NewOpenCell(cells [][]interactor.CellDTO, pos shared.Position) OpenCell {
	return OpenCell{nil, cells[pos.Y][pos.X].BombCount}
}

// そもそもSolved以外(BombCount!=0)のCellを集めているので全てのOpenCellがマッチするのでは？
// -> そのうちCell.Checkedとか実装したら活きてくる。実装してからコメントを外す
//func isSolvedCell(cells [][]interactor.CellDTO, pos shared.Position) bool {
//	// 周りのセルのうち、Cells内に含まれているセルが全てOpenの場合はSolved
//	return pos.ForEachDirectionSatisfy(func(pos shared.Position) (ok bool) {
//		// posがCells内に含まれない場合はSolveに関係ないためok判定
//		return !isInCells(cells, pos) || cells[pos.Y][pos.X].IsOpened
//	})
//}
//
//func isInCells(cells [][]interactor.CellDTO, pos shared.Position) bool {
//	min := 0
//	max := len(cells)
//	return min <= pos.X && pos.X < max && min <= pos.Y && pos.Y < max
//}

package domain

import (
	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

type OpenCell struct {
	shadyPositions []shared.Position
	bombCount      int
	//isSolved  bool
}

func (o OpenCell) GetShadyPositions() []shared.Position {
	return o.shadyPositions
}
func (o OpenCell) GetBombCount() int { return o.bombCount }

func NewSolverCells(cells [][]interactor.CellDTO, bombCounts [][]int) map[shared.Position]OpenCell {
	openCells := make(map[shared.Position]OpenCell)
	for i := range cells {
		for j, cell := range cells[i] {
			if cell.IsOpened && bombCounts[i][j] != 0 {
				pos := shared.NewPosition(j, i)
				openCells[pos] = OpenCell{
					shadyPositions: findShadyPositionsFrom(cells, pos),
					bombCount:      bombCounts[pos.Y][pos.X],
				}
			}
		}
	}
	return openCells
}

func findShadyPositionsFrom(cells [][]interactor.CellDTO, pos shared.Position) []shared.Position {
	shadyPositions := make([]shared.Position, 0, 8)
	pos.ForEachNeighbor(func(pos shared.Position) {
		if !isInCells(cells, pos) {
			return
		}
		cell := cells[pos.Y][pos.X]
		if !cell.IsOpened && !cell.IsFlagged {
			shadyPositions = append(shadyPositions, pos)
		}
	})
	return shadyPositions
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
	max := len(cells)
	return pos.IsInside(max, max)
}

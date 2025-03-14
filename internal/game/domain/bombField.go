package domain

import (
	"container/list"
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type BombField struct {
	board           *board
	closedCellCount int
	checkedCellMap  map[shared.Position]struct{}
	bombCounts      [][]int
	totalBomb       int
}

func (bf *BombField) GetCells() [][]Cell                              { return bf.board.GetCells() }
func (bf *BombField) GetCheckedCellMap() map[shared.Position]struct{} { return bf.checkedCellMap }
func (bf *BombField) GetBombCounts() [][]int                          { return bf.bombCounts }
func (bf *BombField) IsPeaceFul() bool                                { return bf.closedCellCount == bf.totalBomb }

func NewBombField(width int, totalBomb int) *BombField {
	return &BombField{
		board:           NewBoard(width),
		closedCellCount: width * width,
		checkedCellMap:  make(map[shared.Position]struct{}, totalBomb),
		bombCounts:      initBombCounts(width),
		totalBomb:       totalBomb,
	}
}

func initBombCounts(width int) [][]int {
	totalBombs := make([][]int, width)
	for i := range width {
		totalBombs[i] = make([]int, width)
		for j := range width {
			totalBombs[i][j] = 0
		}
	}
	return totalBombs
}

func (bf *BombField) SetBombs(positions map[shared.Position]struct{}) error {
	for pos := range positions {
		if !bf.board.contains(pos) {
			return fmt.Errorf("ボード外のポジションは指定できません")
		}
		bf.board.setCellAt(pos, NewBombCell())
		bf.bombCounts[pos.Y][pos.X] = -1
		bf.incrementBombCountForEachNeighbor(pos)
	}
	return nil
}

func (bf *BombField) incrementBombCountForEachNeighbor(pos shared.Position) {
	pos.ForEachNeighbor(func(p shared.Position) {
		if bf.board.contains(p) && bf.bombCounts[p.Y][p.X] != -1 {
			bf.bombCounts[p.Y][p.X]++
		}
	})
}

func (bf *BombField) OpenCell(pos shared.Position) (bursted bool, err error) {
	bursted, err = bf.openCell(pos)
	if err != nil {
		return false, err
	}
	if bursted {
		return true, nil
	}
	bf.expandSafeArea(pos)
	return false, nil
}

func (bf *BombField) openCell(pos shared.Position) (bursted bool, err error) {
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return false, err
	}
	openedCell, err := cell.WithOpen()
	if err != nil {
		return false, err
	}
	bf.board.setCellAt(pos, openedCell)
	bf.closedCellCount--
	return openedCell.IsBomb(), nil
}

func (bf *BombField) expandSafeArea(pos shared.Position) {
	visited := make([][]bool, bf.board.GetWidth())
	for i := range visited {
		visited[i] = make([]bool, bf.board.GetWidth())
	}
	queue := list.New()
	queue.PushBack(pos)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		pos := front.Value.(shared.Position)
		visited[pos.Y][pos.X] = true
		_, _ = bf.openCell(pos)
		if bf.bombCounts[pos.Y][pos.X] == 0 {
			pos.ForEachNeighbor(func(p shared.Position) {
				cell, err := bf.board.GetCellAt(p)
				if err == nil && !visited[p.Y][p.X] && !cell.IsOpened() {
					queue.PushBack(p)
				}
			})
		}
	}
}

func (bf *BombField) CheckCell(pos shared.Position) error {
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return err
	}
	if cell.isOpened {
		return fmt.Errorf("開放済みのセルです")
	}
	bf.checkedCellMap[pos] = struct{}{}
	return nil
}

func (bf *BombField) UnCheckCell(pos shared.Position) error {
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return err
	}
	if cell.isOpened {
		return fmt.Errorf("開放済みのセルです")
	}
	delete(bf.checkedCellMap, pos)
	return nil
}

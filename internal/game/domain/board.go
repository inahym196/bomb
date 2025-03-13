package domain

import (
	"container/list"
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type board struct {
	width int
	cells [][]Cell
}

func (b *board) GetWidth() int      { return b.width }
func (b *board) GetCells() [][]Cell { return b.cells }
func (b *board) GetCellAt(pos shared.Position) (*Cell, error) {
	if !b.contains(pos) {
		return nil, fmt.Errorf("不正なポジションが選択されました. 有効なrow, columnの範囲は[0-%d]です", b.width-1)
	}
	return &b.cells[pos.Y][pos.X], nil
}

func (b *board) contains(pos shared.Position) bool        { return pos.IsInside(b.width, b.width) }
func (b *board) setCellAt(pos shared.Position, cell Cell) { b.cells[pos.Y][pos.X] = cell }

type BombField struct {
	board           *board
	closedCellCount int
	checkedCellMap  map[shared.Position]struct{}
	bombCounts      [][]int
}

func (bf *BombField) GetWidth() int                                   { return bf.board.GetWidth() }
func (bf *BombField) GetCells() [][]Cell                              { return bf.board.GetCells() }
func (bf *BombField) GetClosedCellCount() (count int)                 { return bf.closedCellCount }
func (bf *BombField) GetCheckedCellMap() map[shared.Position]struct{} { return bf.checkedCellMap }
func (bf *BombField) GetBombCounts() [][]int                          { return bf.bombCounts }
func (bf *BombField) GetCellAt(pos shared.Position) (*Cell, error)    { return bf.board.GetCellAt(pos) }

func (bf *BombField) contains(pos shared.Position) bool        { return bf.board.contains(pos) }
func (bf *BombField) setCellAt(pos shared.Position, cell Cell) { bf.board.setCellAt(pos, cell) }

func NewBombField(width int) *BombField {
	b := &board{width: width, cells: initCells(width)}
	return &BombField{b, width * width, map[shared.Position]struct{}{}, initBombCounts(width)}
}

func initCells(width int) [][]Cell {
	cells := make([][]Cell, width)
	for i := range width {
		cells[i] = make([]Cell, width)
		for j := range width {
			cells[i][j] = NewSafeCell()
		}
	}
	return cells
}

func initBombCounts(width int) [][]int {
	bombCounts := make([][]int, width)
	for i := range width {
		bombCounts[i] = make([]int, width)
		for j := range width {
			bombCounts[i][j] = 0
		}
	}
	return bombCounts
}

func (bf *BombField) SetBombs(positions map[shared.Position]struct{}) error {
	for pos := range positions {
		if !bf.contains(pos) {
			return fmt.Errorf("ボード外のポジションは指定できません")
		}
		bf.setCellAt(pos, NewBombCell())
		bf.bombCounts[pos.Y][pos.X] = -1
		bf.incrementBombCountForEachNeighbor(pos)
	}
	return nil
}

func (bf *BombField) incrementBombCountForEachNeighbor(pos shared.Position) {
	pos.ForEachNeighbor(func(p shared.Position) {
		if bf.contains(p) && bf.bombCounts[p.Y][p.X] != -1 {
			bf.bombCounts[p.Y][p.X]++
		}
	})
}

func (bf *BombField) OpenCell(pos shared.Position) error {
	if err := bf.openCell(pos); err != nil {
		return err
	}
	bf.expandOpenArea(pos)
	return nil
}

func (bf *BombField) openCell(pos shared.Position) error {
	cell, err := bf.GetCellAt(pos)
	if err != nil {
		return err
	}
	openedCell, err := cell.WithOpen()
	if err != nil {
		return err
	}
	bf.setCellAt(pos, openedCell)
	bf.closedCellCount--
	return nil
}

func (bf *BombField) expandOpenArea(pos shared.Position) {
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
		_ = bf.openCell(pos)
		if bf.bombCounts[pos.Y][pos.X] == 0 {
			pos.ForEachNeighbor(func(p shared.Position) {
				cell, err := bf.GetCellAt(p)
				if err == nil && !visited[p.Y][p.X] && !cell.IsOpened() {
					queue.PushBack(p)
				}
			})
		}
	}
}

func (bf *BombField) CheckCell(pos shared.Position) error {
	cell, err := bf.GetCellAt(pos)
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
	cell, err := bf.GetCellAt(pos)
	if err != nil {
		return err
	}
	if cell.isOpened {
		return fmt.Errorf("開放済みのセルです")
	}
	delete(bf.checkedCellMap, pos)
	return nil
}

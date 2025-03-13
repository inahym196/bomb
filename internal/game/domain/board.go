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

type Board struct {
	board           *board
	closedCellCount int
	checkedCellMap  map[shared.Position]struct{}
	bombCounts      [][]int
}

func (b *Board) GetWidth() int                                   { return b.board.GetWidth() }
func (b *Board) GetCells() [][]Cell                              { return b.board.GetCells() }
func (b *Board) GetClosedCellCount() (count int)                 { return b.closedCellCount }
func (b *Board) GetCheckedCellMap() map[shared.Position]struct{} { return b.checkedCellMap }
func (b *Board) GetBombCounts() [][]int                          { return b.bombCounts }
func (b *Board) GetCellAt(pos shared.Position) (*Cell, error)    { return b.board.GetCellAt(pos) }

func (b *Board) contains(pos shared.Position) bool        { return b.board.contains(pos) }
func (b *Board) setCellAt(pos shared.Position, cell Cell) { b.board.setCellAt(pos, cell) }

func NewBoard(width int) *Board {
	b := &board{width: width, cells: initCells(width)}
	return &Board{b, width * width, map[shared.Position]struct{}{}, initBombCounts(width)}
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

func (b *Board) SetBombs(positions map[shared.Position]struct{}) error {
	for pos := range positions {
		if !b.contains(pos) {
			return fmt.Errorf("ボード外のポジションは指定できません")
		}
		b.setCellAt(pos, NewBombCell())
		b.bombCounts[pos.Y][pos.X] = -1
		b.incrementBombCountForEachNeighbor(pos)
	}
	return nil
}

func (b *Board) incrementBombCountForEachNeighbor(pos shared.Position) {
	pos.ForEachNeighbor(func(p shared.Position) {
		if b.contains(p) && b.bombCounts[p.Y][p.X] != -1 {
			b.bombCounts[p.Y][p.X]++
		}
	})
}

func (b *Board) OpenCell(pos shared.Position) error {
	if err := b.openCell(pos); err != nil {
		return err
	}
	b.expandOpenArea(pos)
	return nil
}

func (b *Board) openCell(pos shared.Position) error {
	cell, err := b.GetCellAt(pos)
	if err != nil {
		return err
	}
	openedCell, err := cell.WithOpen()
	if err != nil {
		return err
	}
	b.setCellAt(pos, openedCell)
	b.closedCellCount--
	return nil
}

func (b *Board) expandOpenArea(pos shared.Position) {
	visited := make([][]bool, b.board.GetWidth())
	for i := range visited {
		visited[i] = make([]bool, b.board.GetWidth())
	}
	queue := list.New()
	queue.PushBack(pos)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		pos := front.Value.(shared.Position)
		visited[pos.Y][pos.X] = true
		_ = b.openCell(pos)
		if b.bombCounts[pos.Y][pos.X] == 0 {
			pos.ForEachNeighbor(func(p shared.Position) {
				cell, err := b.GetCellAt(p)
				if err == nil && !visited[p.Y][p.X] && !cell.IsOpened() {
					queue.PushBack(p)
				}
			})
		}
	}
}

func (b *Board) CheckCell(pos shared.Position) error {
	cell, err := b.GetCellAt(pos)
	if err != nil {
		return err
	}
	if cell.isOpened {
		return fmt.Errorf("開放済みのセルです")
	}
	b.checkedCellMap[pos] = struct{}{}
	return nil
}

func (b *Board) UnCheckCell(pos shared.Position) error {
	cell, err := b.GetCellAt(pos)
	if err != nil {
		return err
	}
	if cell.isOpened {
		return fmt.Errorf("開放済みのセルです")
	}
	delete(b.checkedCellMap, pos)
	return nil
}

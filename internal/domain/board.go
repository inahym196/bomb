package domain

import (
	"container/list"
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type Board struct {
	width           int
	cells           [][]Cell
	closedCellCount int
}

func (b *Board) GetCells() [][]Cell                     { return b.cells }
func (b *Board) GetWidth() int                          { return b.width }
func (b *Board) MustGetCellAt(pos shared.Position) Cell { return b.cells[pos.Y][pos.X] }
func (b *Board) GetClosedCellCount() (count int)        { return b.closedCellCount }
func (b *Board) GetCellAt(pos shared.Position) (Cell, error) {
	if !b.inBoard(pos) {
		return NewCell(false), fmt.Errorf("不正な範囲が選択されました. 有効なrowは[0-%d], columnは[0-%d]です", b.width-1, b.width-1)
	}
	return b.cells[pos.Y][pos.X], nil
}

func (b *Board) setCellAt(pos shared.Position, cell Cell) { b.cells[pos.Y][pos.X] = cell }

func NewBoard(width int) *Board {
	return &Board{width, initCells(width), width * width}
}

func initCells(width int) [][]Cell {
	cells := make([][]Cell, width)
	for i := range width {
		cells[i] = make([]Cell, width)
		for j := range width {
			cells[i][j] = NewCell(false)
		}
	}
	return cells
}

func (b *Board) SetBomb(pos shared.Position) {
	b.setCellAt(pos, NewCell(true))
}

func (b *Board) OpenCell(pos shared.Position) error {
	cell, err := b.GetCellAt(pos)
	if err != nil {
		return err
	}
	openedCell, err := cell.Open()
	if err != nil {
		return err
	}
	b.setCellAt(pos, openedCell)
	b.closedCellCount--
	b.expandOpenArea(pos)
	return nil
}

func (b *Board) expandOpenArea(pos shared.Position) {
	visited := map[shared.Position]bool{}
	queue := list.New()
	queue.PushBack(pos)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		pos := front.Value.(shared.Position)
		cell, err := b.GetCellAt(pos)
		if err != nil {
			continue
		}
		visited[pos] = true
		if openedCell, err := cell.Open(); err == nil {
			b.setCellAt(pos, openedCell)
			b.closedCellCount--
		}
		if cell.bombCount == 0 {
			pos.ForEachDirection(func(p shared.Position) {
				if b.inBoard(p) && !visited[p] && !b.MustGetCellAt(p).IsOpened() {
					queue.PushBack(p)
				}
			})
		}
	}
}

func (b *Board) inBoard(pos shared.Position) bool {
	return 0 <= pos.Y && pos.Y < b.width && 0 <= pos.X && pos.X < b.width
}

func (b *Board) IncrementBombCount(pos shared.Position) (err error) {
	cell, err := b.GetCellAt(pos)
	if err != nil {
		return err
	}
	newCell, err := cell.IncrementBombCount()
	if err != nil {
		return err
	}
	b.setCellAt(pos, newCell)
	return nil
}

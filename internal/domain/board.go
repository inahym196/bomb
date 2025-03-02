package domain

import (
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type Board struct {
	width           int
	cells           [][]Cell
	closedCellCount int
}

func (b *Board) GetCells() [][]Cell                 { return b.cells }
func (b *Board) GetCellAt(pos shared.Position) Cell { return b.cells[pos.Y][pos.X] }
func (b *Board) GetClosedCellCount() (count int)    { return b.closedCellCount }

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
	b.cells[pos.Y][pos.X] = NewCell(true)
}

func (b *Board) OpenCell(pos shared.Position) error {
	if !b.inBoard(pos) {
		return fmt.Errorf("不正な範囲が選択されました. 有効なrowは[0-%d], columnは[0-%d]です", b.width-1, b.width-1)
	}
	openedCell, err := b.cells[pos.Y][pos.X].Open()
	if err != nil {
		return err
	}
	b.cells[pos.Y][pos.X] = openedCell
	b.closedCellCount--
	return nil
}

func (b *Board) inBoard(pos shared.Position) bool {
	return 0 <= pos.Y && pos.Y < b.width && 0 <= pos.X && pos.X < b.width
}

func (b *Board) IncrementBombCount(pos shared.Position) (err error) {
	if b.cells[pos.Y][pos.X], err = b.cells[pos.Y][pos.X].IncrementBombCount(); err != nil {
		return err
	}
	return nil
}

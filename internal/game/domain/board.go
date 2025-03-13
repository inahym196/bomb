package domain

import (
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

func NewBoard(width int) *board {
	return &board{width: width, cells: initCells(width)}
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

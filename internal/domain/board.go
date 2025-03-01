package domain

import "fmt"

type Cell struct {
	isOpened bool
	isBomb   bool
}

func (c Cell) IsOpened() bool { return c.isOpened }
func (c Cell) IsBomb() bool   { return c.isBomb }

func NewCell(isBomb bool) Cell {
	return Cell{false, isBomb}
}

func (c Cell) Open() (Cell, error) {
	if c.isOpened {
		return Cell{}, fmt.Errorf("すでに開放済みのセルです")
	}
	return Cell{isOpened: true, isBomb: c.isBomb}, nil
}

type Board struct {
	width int
	cells [][]Cell
}

func (b *Board) GetCells() [][]Cell { return b.cells }

func NewBoard(width int) *Board {
	return &Board{width, initCells(width)}
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

func (b *Board) SetBomb(row, col int) {
	b.cells[row][col] = NewCell(true)
}

func (b *Board) OpenCell(row, col int) error {
	if !b.inBoard(row, col) {
		return fmt.Errorf("不正な範囲が選択されました. 有効なrowは[0-%d], columnは[0-%d]です", b.width-1, b.width-1)
	}
	openedCell, err := b.cells[row][col].Open()
	if err != nil {
		return err
	}
	b.cells[row][col] = openedCell
	return nil
}

func (b *Board) inBoard(row, col int) bool {
	return 0 <= row && row < b.width && 0 <= col && col < b.width
}

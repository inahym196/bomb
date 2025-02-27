package domain

import "fmt"

type Cell struct {
	isOpened bool
	isBomb   bool
}

func (c Cell) IsOpened() bool { return c.isOpened }
func (c Cell) IsBomb() bool   { return c.isBomb }

func NewCell(isOpened, isBomb bool) Cell {
	return Cell{isOpened, isBomb}
}

func (c *Cell) Open() error {
	if c.isOpened {
		return fmt.Errorf("すでに開放済みのセルです")
	}
	c.isOpened = true
	return nil
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
			cells[i][j] = NewCell(false, false)
		}
	}
	return cells
}

type Position struct {
	Row int
	Col int
}

func (b *Board) SetBombs(positionList []Position) {
	for _, pos := range positionList {
		b.cells[pos.Row][pos.Col] = NewCell(false, true)
	}
}

func (b *Board) inBoard(row, col int) bool {
	return 0 <= row && row < b.width && 0 <= col && col < b.width
}

func (b *Board) OpenCell(row, col int) error {
	if !b.inBoard(row, col) {
		return fmt.Errorf("範囲内のセルを選択してください. rowは[0-%d], columnは[0-%d]", b.width-1, b.width-1)
	}
	return b.cells[row][col].Open()
}

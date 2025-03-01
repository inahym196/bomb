package domain

import "fmt"

type Cell struct {
	isOpened  bool
	isBomb    bool
	bombCount int
}

func (c Cell) IsOpened() bool    { return c.isOpened }
func (c Cell) IsBomb() bool      { return c.isBomb }
func (c Cell) GetBombCount() int { return c.bombCount }

func NewCell(isBomb bool) Cell {
	if isBomb {
		return Cell{false, true, -1}
	}
	return Cell{false, false, 0}
}

func (c Cell) Open() (Cell, error) {
	if c.isOpened {
		return Cell{}, fmt.Errorf("すでに開放済みのセルです")
	}
	return Cell{isOpened: true, isBomb: c.isBomb, bombCount: c.bombCount}, nil
}

func (c Cell) IncrementBombCount() (Cell, error) {
	if c.bombCount == -1 {
		return Cell{}, fmt.Errorf("ボムのマスは計算対象外です")
	}
	return Cell{
		isOpened:  c.isOpened,
		isBomb:    c.isBomb,
		bombCount: c.bombCount + 1,
	}, nil
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

func (b *Board) IncrementBombCount(row, col int) (err error) {
	if b.cells[row][col], err = b.cells[row][col].IncrementBombCount(); err != nil {
		return err
	}
	return nil
}

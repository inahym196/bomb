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
		return Cell{}, fmt.Errorf("ボムのセルは計算対象外です")
	}
	return Cell{
		isOpened:  c.isOpened,
		isBomb:    c.isBomb,
		bombCount: c.bombCount + 1,
	}, nil
}

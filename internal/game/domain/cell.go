package domain

import "fmt"

type Cell struct {
	isOpened bool
	isBomb   bool
}

func (c Cell) IsOpened() bool { return c.isOpened }
func (c Cell) IsBomb() bool   { return c.isBomb }

func NewCell(isBomb bool) Cell {
	if isBomb {
		return Cell{false, true}
	}
	return Cell{false, false}
}

func (c Cell) WithOpen() (Cell, error) {
	if c.isOpened {
		return Cell{}, fmt.Errorf("すでに開放済みのセルです")
	}
	return Cell{isOpened: true, isBomb: c.isBomb}, nil
}

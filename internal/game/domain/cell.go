package domain

import (
	"errors"
)

type Cell struct {
	isOpened  bool
	isBomb    bool
	isFlagged bool
}

func (c Cell) IsOpened() bool  { return c.isOpened }
func (c Cell) IsBomb() bool    { return c.isBomb }
func (c Cell) IsFlagged() bool { return c.isFlagged }
func (c Cell) CanOpen() bool   { return !c.isOpened && !c.isFlagged }

func NewSafeCell() Cell {
	return Cell{
		isOpened:  false,
		isBomb:    false,
		isFlagged: false,
	}
}

func NewBombCell() Cell {
	return Cell{
		isOpened:  false,
		isBomb:    true,
		isFlagged: false,
	}
}

func (c Cell) Equals(target Cell) bool {
	return c.isBomb == target.isBomb &&
		c.isFlagged == target.isFlagged &&
		c.isOpened == target.isOpened
}

func (c Cell) Open() (Cell, error) {
	if c.isOpened {
		return Cell{}, errors.New("cell is already opened")
	}
	if c.isFlagged {
		return Cell{}, errors.New("cannot open a flagged cell")
	}
	return Cell{
		isOpened:  true,
		isBomb:    c.isBomb,
		isFlagged: c.isFlagged,
	}, nil
}

func (c Cell) Flagged() (Cell, error) {
	if c.isOpened {
		return Cell{}, errors.New("cell is already opened")
	}
	if c.isFlagged {
		return Cell{}, errors.New("cell is already flagged")
	}
	return Cell{
		isOpened:  c.isOpened,
		isBomb:    c.isBomb,
		isFlagged: true,
	}, nil
}

func (c Cell) UnFlagged() (Cell, error) {
	if c.isOpened {
		return Cell{}, errors.New("cell is already opened")
	}
	if !c.isFlagged {
		return Cell{}, errors.New("cell is already unflagged")
	}
	return Cell{
		isOpened:  c.isOpened,
		isBomb:    c.isBomb,
		isFlagged: false,
	}, nil
}

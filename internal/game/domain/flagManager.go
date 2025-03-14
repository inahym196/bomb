package domain

import (
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type FlagManager struct {
	flagMap [][]bool
}

func (fc *FlagManager) GetFlagMap() [][]bool { return fc.flagMap }

func NewFlagManager(size int) *FlagManager {
	flagMap := make([][]bool, size)
	for i := range size {
		flagMap[i] = make([]bool, size)
	}
	return &FlagManager{flagMap}
}

func (fc *FlagManager) Contains(pos shared.Position) bool {
	n := len(fc.flagMap)
	return pos.IsInside(n, n)
}

func (fc *FlagManager) IsFlagged(pos shared.Position) bool {
	return fc.Contains(pos) && fc.flagMap[pos.Y][pos.X]
}

func (fc *FlagManager) Flag(pos shared.Position, cell Cell) error {
	if cell.IsOpened() {
		return fmt.Errorf("開放済みのセルです")
	}
	if !fc.Contains(pos) {
		// TODO: 到達できないことを確認する
		return fmt.Errorf("想定外")
	}
	fc.flagMap[pos.Y][pos.X] = true
	return nil
}

func (fc *FlagManager) UnFlag(pos shared.Position, cell Cell) error {
	if cell.IsOpened() {
		return fmt.Errorf("開放済みのセルです")
	}
	if !fc.Contains(pos) {
		// TODO: 到達できないことを確認する
		return fmt.Errorf("想定外")
	}
	fc.flagMap[pos.Y][pos.X] = false
	return nil
}

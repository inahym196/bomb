package domain

import (
	"reflect"
	"testing"

	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

var (
	opened  = interactor.CellDTO{IsOpened: true, IsBomb: false, IsFlagged: false}
	closed  = interactor.CellDTO{IsOpened: false, IsBomb: false, IsFlagged: false}
	flagged = interactor.CellDTO{IsOpened: false, IsBomb: false, IsFlagged: true}
)

func BenchmarkNewSolverCells(b *testing.B) {
	cells := [][]interactor.CellDTO{
		{closed, flagged, closed},
		{opened, opened, opened},
		{opened, opened, opened},
	}
	bombCounts := [][]int{
		{1, -1, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	for i := range cells {
		for j := range cells[i] {
			cells[i][j].BombCount = bombCounts[i][j]
		}
	}
	b.ResetTimer()
	for range b.N {
		NewSolverCells(cells)
	}
}

func TestNewSolverCells(t *testing.T) {
	cells := [][]interactor.CellDTO{
		{closed, flagged, closed},
		{opened, opened, opened},
		{opened, opened, opened},
	}
	bombCounts := [][]int{
		{1, -1, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	for i := range cells {
		for j := range cells[i] {
			cells[i][j].BombCount = bombCounts[i][j]
		}
	}
	want := map[shared.Position]OpenCell{
		shared.NewPosition(0, 1): {
			[]shared.Position{shared.NewPosition(0, 0)}, 1,
		},
		shared.NewPosition(1, 1): {
			[]shared.Position{shared.NewPosition(0, 0), shared.NewPosition(2, 0)}, 1,
		},
		shared.NewPosition(2, 1): {
			[]shared.Position{shared.NewPosition(2, 0)}, 1,
		},
	}
	if got := NewSolverCells(cells); !reflect.DeepEqual(got, want) {
		t.Errorf("NewSolverCells() = %v, want %v", got, want)
	}
}

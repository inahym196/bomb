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
		{closed, closed, closed},
		{opened, flagged, opened},
		{opened, opened, opened},
	}
	bombCounts := [][]int{
		{1, -1, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	b.ResetTimer()
	for range b.N {
		NewSolverCells(cells, bombCounts)
	}
}

func TestNewSolverCells(t *testing.T) {
	cells := [][]interactor.CellDTO{
		{closed, closed, closed},
		{opened, flagged, opened},
		{opened, opened, opened},
	}
	bombCounts := [][]int{
		{1, -1, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	want := map[shared.Position]OpenCell{
		shared.NewPosition(0, 1): {
			map[shared.Position]struct{}{shared.NewPosition(0, 0): {}, shared.NewPosition(1, 0): {}}, 1,
		},
		shared.NewPosition(2, 1): {
			map[shared.Position]struct{}{shared.NewPosition(1, 0): {}, shared.NewPosition(2, 0): {}}, 1,
		},
	}
	if got := NewSolverCells(cells, bombCounts); !reflect.DeepEqual(got, want) {
		t.Errorf("NewSolverCells() = %v, want %v", got, want)
	}
}

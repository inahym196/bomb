package domain

import (
	"reflect"
	"testing"

	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

func Test_theorem3_Apply(t *testing.T) {
	closed := interactor.CellDTO{IsOpened: false, IsBomb: false, IsFlagged: false}
	opened := interactor.CellDTO{IsOpened: true, IsBomb: false, IsFlagged: false}
	bcs := [][]int{
		{0, 1, 0},
		{0, 3, 0},
		{0, 0, 0},
	}
	cells := [][]interactor.CellDTO{
		{closed, opened, closed},
		{closed, opened, closed},
		{closed, opened, closed},
	}
	for i := range bcs {
		for j := range bcs[i] {
			cells[i][j].BombCount = bcs[i][j]
		}
	}
	want := Solution{[]shared.Position{{X: 0, Y: 2}, {X: 2, Y: 2}}, SolutionResultIsBomb}
	tr := theorem3{}
	if got := tr.Apply(cells); !reflect.DeepEqual(got, want) {
		t.Errorf("theorem3.Apply().Positions = %v, want %v", got.Positions, want.Positions)
	}
}

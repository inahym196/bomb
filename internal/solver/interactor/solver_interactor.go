package interactor

import (
	"github.com/inahym196/bomb/internal/interactor"
	"github.com/inahym196/bomb/internal/solver/domain"
)

type GameSolverInteractor struct{}

func NewGameSolverInteractor() *GameSolverInteractor {
	return &GameSolverInteractor{}
}

type GetHintParam struct {
	Cells [][]interactor.CellDTO
}

func (si *GameSolverInteractor) GetHint(param GetHintParam) string {
	cells := domain.NewSolverCells(param.Cells)
	return domain.NewSolver(cells).Solve()
}

package interactor

import (
	"github.com/inahym196/bomb/internal/game/interactor"
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
	return domain.NewSolver(param.Cells).Solve()
}

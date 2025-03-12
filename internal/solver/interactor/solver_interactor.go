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
	Game interactor.GameDTO
}

func (si *GameSolverInteractor) GetHint(param GetHintParam) string {
	cells := domain.NewSolverCells(param.Game)
	return domain.NewSolver(cells).Solve()
}

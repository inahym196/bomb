package interactor

import "github.com/inahym196/bomb/internal/domain"

type GameInteractor struct {
	game domain.Game
}

func NewGameInteractor() *GameInteractor {
	opt := &domain.GameOption{
		BoardWidth: 8,
		Bombs:      []domain.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	}
	return &GameInteractor{*domain.NewGame(opt)}
}

func (gi *GameInteractor) GetBoardCellStates() GetBoardCellStatesOutput {
	cells := gi.game.GetBoard().GetCells()
	return GetBoardCellStatesOutput{
		CellStates: CellStatesFrom(cells),
	}
}

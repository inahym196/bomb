package interactor

import (
	"fmt"

	"github.com/inahym196/bomb/internal/domain"
)

var ImmemoryGame *domain.Game

type GameInteractor struct{}

func NewGameInteractor() *GameInteractor {
	return &GameInteractor{}
}

type Position struct {
	Row int
	Col int
}

type InitGameParam struct {
	BoardWidth int
	Bombs      []Position
}
type InitGameResult struct {
	GameDTO
}

func (gi *GameInteractor) InitGame(param InitGameParam) *InitGameResult {
	opt := &domain.GameOption{
		BoardWidth: param.BoardWidth,
		Bombs:      toPositions(param.Bombs),
	}
	ImmemoryGame = domain.NewGame(opt)
	dto := toGameDTO(ImmemoryGame)
	return &InitGameResult{dto}
}

func toPositions(dto []Position) []domain.Position {
	poss := make([]domain.Position, len(dto))
	for i, pos := range dto {
		poss[i] = domain.Position{Row: pos.Row, Col: pos.Col}
	}
	return poss
}

type GetGameResult struct {
	GameDTO
}

func (gi *GameInteractor) GetGame() (GetGameResult, error) {
	game := ImmemoryGame
	if game == nil {
		return GetGameResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	return GetGameResult{toGameDTO(game)}, nil
}

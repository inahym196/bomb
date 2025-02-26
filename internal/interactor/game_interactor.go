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

type InitGameParam struct {
	BoardWidth int
	BombCount  int
}
type InitGameResult struct {
	GameDTO
}

func (gi *GameInteractor) InitGame(param InitGameParam) *InitGameResult {
	opt := &domain.GameOption{
		BoardWidth: param.BoardWidth,
		BombCount:  param.BombCount,
	}
	ImmemoryGame = domain.NewGame(opt)
	dto := toGameDTO(ImmemoryGame)
	return &InitGameResult{dto}
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

type OpenCellParam struct {
	Row int
	Col int
}

type OpenCellResult struct {
	GameDTO
}

func (gi *GameInteractor) OpenCell(OpenCellParam) (OpenCellResult, error) {
	game := ImmemoryGame
	if game == nil {
		return OpenCellResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	return OpenCellResult{toGameDTO(game)}, nil
}

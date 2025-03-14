package interactor

import (
	"fmt"

	"github.com/inahym196/bomb/internal/game/domain"
	"github.com/inahym196/bomb/pkg/shared"
)

var ImmemoryGame *domain.Game

type ImmemoryGameRepository struct {
	game *domain.Game
}

func NewInmemoryGameRepository() *ImmemoryGameRepository {
	return &ImmemoryGameRepository{ImmemoryGame}
}

func (igr *ImmemoryGameRepository) Find() (game *domain.Game, ok bool) {
	if igr.game == nil {
		return nil, false
	}
	return igr.game, true
}

func (igr *ImmemoryGameRepository) Save(game *domain.Game) error {
	if game == nil {
		return fmt.Errorf("nil pointer error")
	}
	igr.game = game
	return nil
}

type GameInteractor struct {
	game_repo *ImmemoryGameRepository
}

func NewGameInteractor() *GameInteractor {
	return &GameInteractor{NewInmemoryGameRepository()}
}

type InitGameParam struct {
	BoardWidth int
	BombCount  int
}
type InitGameResult struct {
	GameDTO
}

func (gi *GameInteractor) InitGame(param InitGameParam) (InitGameResult, error) {
	newGame, err := domain.NewGame(param.BoardWidth, param.BombCount)
	if err != nil {
		return InitGameResult{}, err
	}
	if err := gi.game_repo.Save(newGame); err != nil {
		return InitGameResult{}, err
	}
	return InitGameResult{toGameDTO(newGame)}, nil
}

type GetGameResult struct {
	GameDTO
}

func (gi *GameInteractor) GetGame() (GetGameResult, error) {
	game, ok := gi.game_repo.Find()
	if !ok {
		return GetGameResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	return GetGameResult{toGameDTO(game)}, nil
}

type OpenCellParam struct {
	Pos shared.Position
}

type OpenCellResult struct {
	GameDTO
}

func (gi *GameInteractor) OpenCell(param OpenCellParam) (OpenCellResult, error) {
	game, ok := gi.game_repo.Find()
	if !ok {
		return OpenCellResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	if err := game.OpenCell(param.Pos); err != nil {
		return OpenCellResult{}, err
	}
	if err := gi.game_repo.Save(game); err != nil {
		return OpenCellResult{}, err
	}
	return OpenCellResult{toGameDTO(game)}, nil
}

type CheckCellParam struct {
	Pos shared.Position
}

type CheckCellResult struct {
	GameDTO
}

func (gi *GameInteractor) CheckCell(param CheckCellParam) (CheckCellResult, error) {
	game, ok := gi.game_repo.Find()
	if !ok {
		return CheckCellResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	if err := game.Flag(param.Pos); err != nil {
		return CheckCellResult{}, err
	}
	if err := gi.game_repo.Save(game); err != nil {
		return CheckCellResult{}, err
	}
	return CheckCellResult{toGameDTO(game)}, nil
}

type UnCheckCellParam struct {
	Pos shared.Position
}

type UnCheckCellResult struct {
	GameDTO
}

func (gi *GameInteractor) UnCheckCell(param UnCheckCellParam) (UnCheckCellResult, error) {
	game, ok := gi.game_repo.Find()
	if !ok {
		return UnCheckCellResult{}, fmt.Errorf("ゲームが初期化されていません")
	}
	if err := game.UnFlag(param.Pos); err != nil {
		return UnCheckCellResult{}, err
	}
	if err := gi.game_repo.Save(game); err != nil {
		return UnCheckCellResult{}, err
	}
	return UnCheckCellResult{toGameDTO(game)}, nil
}

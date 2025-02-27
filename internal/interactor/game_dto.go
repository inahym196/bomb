package interactor

import (
	"github.com/inahym196/bomb/internal/domain"
)

type CellDTO struct {
	IsOpened bool
	IsBomb   bool
}

func CellsFrom(cells [][]domain.Cell) [][]CellDTO {
	dto := make([][]CellDTO, len(cells))
	for i, row := range cells {
		dto[i] = make([]CellDTO, len(row))
		for j, cell := range row {
			dto[i][j] = CellFrom(cell)
		}
	}
	return dto
}

func CellFrom(cell domain.Cell) CellDTO {
	return CellDTO{
		IsOpened: cell.IsOpened(),
		IsBomb:   cell.IsBomb(),
	}
}

const (
	GameStateReady byte = iota
	GameStatePlaying
	GameStateCompleted
	GameStateFailed
)

type GameDTO struct {
	BoardCells [][]CellDTO
	State      byte
}

func toGameDTO(game *domain.Game) GameDTO {
	return GameDTO{CellsFrom(game.GetBoard().GetCells()), game.GetState()}
}

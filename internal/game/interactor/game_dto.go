package interactor

import (
	"github.com/inahym196/bomb/internal/game/domain"
)

type CellDTO struct {
	IsOpened bool
	IsBomb   bool
}

type GameDTO struct {
	BoardCells [][]CellDTO
	State      byte
	BombCounts [][]int
	FlagMap    [][]bool
}

func toGameDTO(game *domain.Game) GameDTO {
	bombField := game.GetBombField()
	return GameDTO{
		BoardCells: cellsFrom(bombField.GetCells()),
		State:      game.GetState(),
		BombCounts: bombField.GetBombCounts(),
		FlagMap:    bombField.GetFlagMap(),
	}
}

func cellsFrom(cells [][]domain.Cell) [][]CellDTO {
	dto := make([][]CellDTO, len(cells))
	for i, row := range cells {
		dto[i] = make([]CellDTO, len(row))
		for j, cell := range row {
			dto[i][j] = cellFrom(cell)
		}
	}
	return dto
}

func cellFrom(cell domain.Cell) CellDTO {
	return CellDTO{
		IsOpened: cell.IsOpened(),
		IsBomb:   cell.IsBomb(),
	}
}

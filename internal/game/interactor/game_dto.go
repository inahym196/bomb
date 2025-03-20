package interactor

import (
	"github.com/inahym196/bomb/internal/game/domain"
)

type CellDTO struct {
	IsOpened  bool
	IsBomb    bool
	IsFlagged bool
	BombCount int
}

type GameDTO struct {
	BoardCells [][]CellDTO
	State      byte
}

func toGameDTO(game *domain.Game) GameDTO {
	bombField := game.GetBombField()
	return GameDTO{
		BoardCells: cellsFrom(bombField.GetCells(), bombField.GetBombCounts()),
		State:      game.GetState(),
	}
}

func cellsFrom(cells [][]domain.Cell, bombCounts [][]int) [][]CellDTO {
	dto := make([][]CellDTO, len(cells))
	for i, row := range cells {
		dto[i] = make([]CellDTO, len(row))
		for j, cell := range row {
			dto[i][j] = cellFrom(cell, bombCounts[i][j])
		}
	}
	return dto
}

func cellFrom(cell domain.Cell, bombCount int) CellDTO {
	return CellDTO{
		IsOpened:  cell.IsOpened(),
		IsBomb:    cell.IsBomb(),
		IsFlagged: cell.IsFlagged(),
		BombCount: bombCount,
	}
}

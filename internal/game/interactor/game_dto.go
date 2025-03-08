package interactor

import (
	"github.com/inahym196/bomb/internal/game/domain"
	"github.com/inahym196/bomb/pkg/shared"
)

type CellDTO struct {
	IsOpened  bool
	IsBomb    bool
	BombCount int
}

type GameDTO struct {
	BoardCells     [][]CellDTO
	State          byte
	CheckedCellMap map[shared.Position]struct{}
}

func toGameDTO(game *domain.Game) GameDTO {
	return GameDTO{cellsFrom(game.GetBoard().GetCells()), game.GetState(), game.GetBoard().GetCheckedCellMap()}
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
		IsOpened:  cell.IsOpened(),
		IsBomb:    cell.IsBomb(),
		BombCount: cell.GetBombCount(),
	}
}

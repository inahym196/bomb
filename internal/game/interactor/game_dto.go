package interactor

import (
	"github.com/inahym196/bomb/internal/game/domain"
	"github.com/inahym196/bomb/pkg/shared"
)

type CellDTO struct {
	IsOpened bool
	IsBomb   bool
}

type GameDTO struct {
	BoardCells     [][]CellDTO
	State          byte
	CheckedCellMap map[shared.Position]struct{}
	BombCounts     [][]int
}

func toGameDTO(game *domain.Game) GameDTO {
	board := game.GetBoard()
	return GameDTO{cellsFrom(board.GetCells()), game.GetState(), board.GetCheckedCellMap(), board.GetBombCounts()}
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

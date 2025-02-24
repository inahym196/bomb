package interactor

import (
	"github.com/inahym196/bomb/internal/domain"
)

const (
	CellStateUndefined byte = iota
	CellStateClosed
	CellStateBomb
	CellStateOpen
)

func CellStatesFrom(cells [][]domain.Cell) [][]byte {
	bytes := make([][]byte, len(cells))
	for i, row := range cells {
		bytes[i] = make([]byte, len(row))
		for j, cell := range row {
			bytes[i][j] = CellStateFrom(cell)
		}
	}
	return bytes
}

func CellStateFrom(cell domain.Cell) byte {
	switch cell.GetState() {
	case domain.CellClosed:
		return CellStateClosed
	case domain.CellBomb:
		return CellStateBomb
	case domain.CellOpen:
		return CellStateClosed
	default:
		return CellStateUndefined
	}
}

type GetBoardCellStatesOutput struct {
	CellStates [][]byte
}

type GameDTO struct {
	BoardCellStates [][]byte
}

func toGameDTO(game *domain.Game) GameDTO {
	return GameDTO{CellStatesFrom(game.GetBoard().GetCells())}
}

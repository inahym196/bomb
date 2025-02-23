package controller

import (
	"fmt"

	"github.com/inahym196/bomb"
)

type CLIController struct {
	game *bomb.Game
}

func NewCLIController(game *bomb.Game) *CLIController {
	return &CLIController{game}
}

func (c *CLIController) GetBoard() string {
	var output string
	cells := c.game.GetBoard().GetCells()
	cellsDTO := cellsToDTO(cells)
	for _, row := range cellsDTO {
		for _, cell := range row {
			output += fmt.Sprintf(" %s", cell)
		}
		output += "\n"
	}
	return output
}

func cellsToDTO(cells [][]bomb.Cell) [][]string {
	dto := make([][]string, len(cells))
	for i, row := range cells {
		dto[i] = make([]string, len(cells))
		for j, cell := range row {
			dto[i][j] = cellToDTO(cell)
		}
	}
	return dto
}

func cellToDTO(cell bomb.Cell) string {
	switch cell.GetState() {
	case bomb.CellBomb:
		return "B"
	case bomb.CellOpen:
		return " "
	case bomb.CellClosed:
		return "?"
	default:
		return ""
	}
}

package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb"
)

type CLIController struct {
	game *bomb.Game
}

func NewCLIController(game *bomb.Game) *CLIController {
	return &CLIController{game}
}

func (c *CLIController) getBoard() string {
	var output string
	cells := c.game.GetBoard().GetCells()
	cellsDTO := cellsToStr(cells)
	for _, row := range cellsDTO {
		for _, cell := range row {
			output += fmt.Sprintf(" %s", cell)
		}
		output += "\n"
	}
	return output
}

func cellsToStr(cells [][]bomb.Cell) [][]string {
	dto := make([][]string, len(cells))
	for i, row := range cells {
		dto[i] = make([]string, len(cells))
		for j, cell := range row {
			dto[i][j] = cellToStr(cell)
		}
	}
	return dto
}

func cellToStr(cell bomb.Cell) string {
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

func (c *CLIController) Run() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Enter command> ")
		input.Scan()
		switch input.Text() {
		case "show", "s":
			fmt.Print(c.getBoard())
		case "exit", "quit":
			return
		case "help", "h":
			fmt.Print(heredoc.Doc(`
			Usage:
			  > help              this message
			  > show              show board
			  > open <row> <col>  open cell
			  > exit              end game
			`))
		default:
			fmt.Printf("\"%s\" is invalid command. Use \"help\"\n", input.Text())
		}
	}

}

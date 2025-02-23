package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb/internal/domain"
)

type CLIController struct {
	game *domain.Game
}

func NewCLIController(game *domain.Game) *CLIController {
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

func cellsToStr(cells [][]domain.Cell) [][]string {
	dto := make([][]string, len(cells))
	for i, row := range cells {
		dto[i] = make([]string, len(cells))
		for j, cell := range row {
			dto[i][j] = cellToStr(cell)
		}
	}
	return dto
}

func cellToStr(cell domain.Cell) string {
	switch cell.GetState() {
	case domain.CellBomb:
		return "B"
	case domain.CellOpen:
		return " "
	case domain.CellClosed:
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
		case "exit", "quit", "q":
			fmt.Println("Exiting...")
			return
		case "help", "h":
			fmt.Print(heredoc.Doc(`
			Available Commands:
			  > show              Show board
			  > open <row> <col>  Open cell
			  > help              Show this help message
			  > exit              Exit the program
			`))
		default:
			fmt.Printf("\"%s\" is invalid command. Use \"help\"\n", input.Text())
		}
	}

}

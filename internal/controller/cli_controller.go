package controller

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb/internal/interactor"
)

type CLIController struct {
	gi *interactor.GameInteractor
}

func NewCLIController() *CLIController {
	return &CLIController{
		gi: interactor.NewGameInteractor(),
	}
}

func (c *CLIController) Run() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Enter command> ")
		input.Scan()
		text := input.Text()
		switch text {
		case "show", "s":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Print(err.Error())
			}
			fmt.Print(c.parseGame(result.GameDTO))
		case "start":
			result := c.gi.InitGame(interactor.InitGameParam{
				BoardWidth: 8,
				Bombs:      []interactor.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
			})
			fmt.Print(c.parseGame(result.GameDTO))
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
			fmt.Printf("\"%s\" is invalid command. Use \"help\"\n", text)
		}
	}
}

func (c *CLIController) parseGame(game interactor.GameDTO) (output string) {
	for _, row := range game.BoardCellStates {
		for _, state := range row {
			output += fmt.Sprintf(" %s", cellToStr(state))
		}
		output += "\n"
	}
	return output
}

func cellToStr(cell interactor.CellDTO) string {
	if cell.IsOpened {
		return " "
	}
	return "?"
}

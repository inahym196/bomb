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

func NewCLIController(gi *interactor.GameInteractor) *CLIController {
	return &CLIController{gi}
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
			fmt.Printf("\"%s\" is invalid command. Use \"help\"\n", text)
		}
	}
}

func (c *CLIController) getBoard() string {
	var output string
	states := c.gi.GetBoardCellStates().CellStates
	for _, row := range states {
		for _, state := range row {
			output += fmt.Sprintf(" %s", stateToStr(state))
		}
		output += "\n"
	}
	return output
}

func stateToStr(state byte) string {
	switch state {
	case interactor.CellStateBomb:
		return "B"
	case interactor.CellStateOpen:
		return " "
	case interactor.CellStateClosed:
		return "?"
	default:
		return ""
	}
}

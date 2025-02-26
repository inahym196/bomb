package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		fmt.Print("Enter command> ")
		input.Scan()
		text := input.Text()
		words := strings.Fields(text)
		wordLen := len(words)
		if wordLen == 0 {
			continue
		}
		switch words[0] {
		case "start", "restart", "init":
			result := c.gi.InitGame(interactor.InitGameParam{
				BoardWidth: 9,
				BombCount:  10,
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
			fmt.Printf("\"%s\" is invalid command. See \"help\"\n", text)
		}
		fmt.Println()
	}
}

func (c *CLIController) parseGame(game interactor.GameDTO) (output string) {
	for _, row := range game.BoardCellStates {
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

package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			result, err := c.gi.InitGame(interactor.InitGameParam{
				BoardWidth: 9,
				BombCount:  10,
			})
			if err != nil {
				fmt.Print(err.Error())
			}
			fmt.Print(c.parseGame(result.GameDTO))
		case "show":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Print(err.Error())
			}
			fmt.Print(c.parseGame(result.GameDTO))
		case "debug":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Print(err.Error())
			}
			fmt.Print(c.debugGame(result.GameDTO))
		case "open":
			if wordLen != 3 {
				fmt.Println("invalid num of args. See \"help\"")
				continue
			}
			row, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Println("invalid row. row should be a number.")
				continue
			}
			col, err := strconv.Atoi(words[2])
			if err != nil {
				fmt.Println("invalid column. column should be a number.")
				continue
			}
			result, err := c.gi.OpenCell(interactor.OpenCellParam{Row: row, Col: col})
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Print(c.parseGame(result.GameDTO))
		case "exit", "quit", "q":
			fmt.Println("Exiting...")
			return
		case "help", "h":
			fmt.Print(heredoc.Doc(`
			Available Commands:
			  > start             Start Game
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
	for _, row := range game.BoardCells {
		for _, cell := range row {
			output += fmt.Sprintf(" %s", cellToStr(cell))
		}
		output += "\n"
	}
	return output
}

func cellToStr(cell interactor.CellDTO) string {
	if !cell.IsOpened {
		return "□"
	}
	if cell.IsBomb {
		return "B"
	}
	return " "
}
func (c *CLIController) debugGame(game interactor.GameDTO) (output string) {
	for _, row := range game.BoardCells {
		for _, cell := range row {
			output += fmt.Sprintf(" %s", cellToDebugStr(cell))
		}
		output += "\n"
	}
	return output
}

func cellToDebugStr(cell interactor.CellDTO) string {
	if cell.IsBomb {
		return "B"
	}
	if !cell.IsOpened {
		return "□"
	}
	return " "
}

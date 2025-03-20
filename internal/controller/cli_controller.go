package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb/internal/game/domain"
	"github.com/inahym196/bomb/internal/game/interactor"
	si "github.com/inahym196/bomb/internal/solver/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

type CLIController struct {
	gi *interactor.GameInteractor
	si *si.GameSolverInteractor
}

func NewCLIController() *CLIController {
	return &CLIController{
		gi: interactor.NewGameInteractor(),
		si: si.NewGameSolverInteractor(),
	}
}

func (c *CLIController) Run() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command> ")
		input.Scan()
		text := input.Text()
		words := strings.Fields(text)
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "start", "restart", "init":
			var boardWidth, totalBomb int
			mode, err := c.parseStartGameModeArgs(words)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			switch mode {
			case GameModeEasy:
				boardWidth, totalBomb = 9, 10
			case GameModeNormal:
				boardWidth, totalBomb = 16, 40
			case GameModeCustom:
				boardWidth, totalBomb, err = c.parseStartCustomModeArgs(words)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			result, err := c.gi.InitGame(interactor.InitGameParam{
				BoardWidth: boardWidth,
				BombCount:  totalBomb,
			})
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.parseGame(result.GameDTO))
		case "show":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.parseGame(result.GameDTO))
		case "debug":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.debugGame(result.GameDTO))
		case "hint":
			result, err := c.gi.GetGame()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			param := si.GetHintParam{Cells: result.BoardCells}
			fmt.Println(c.si.GetHint(param))
		case "open":
			row, col, err := c.parseOpenArgs(words)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			param := interactor.OpenCellParam{Pos: shared.NewPosition(col, row)}
			result, err := c.gi.OpenCell(param)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.parseGame(result.GameDTO))
		case "check":
			row, col, err := c.parseCheckArgs(words)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			param := interactor.CheckCellParam{Pos: shared.NewPosition(col, row)}
			result, err := c.gi.CheckCell(param)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.parseGame(result.GameDTO))
		case "uncheck":
			row, col, err := c.parseUnCheckArgs(words)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			param := interactor.UnCheckCellParam{Pos: shared.NewPosition(col, row)}
			result, err := c.gi.UnCheckCell(param)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("gameState: %s\n", stateToStr(result.State))
			fmt.Println(c.parseGame(result.GameDTO))
		case "exit", "quit", "q":
			fmt.Println("終了中...")
			return
		case "help", "h":
			fmt.Println(heredoc.Doc(`
			Available Commands:
			  > start <mode: easy or normal>        Start Game, Select gameMode
			  > start custom <width> <totalBomb>    Start Game, Set Custom width and totalBomb
			  > show                                Show board
			  > open <row: int> <col: alpha>        Open cell
			  > check/uncheck <row> <col>           Check/UnCheck cell
			  > help                                Show this help message
			  > exit                                Exit the program
			`))
		default:
			fmt.Printf("\"%s\"は無効なコマンドです. \"help\"コマンドを確認してください\n\n", text)
		}
	}
}

func (c *CLIController) parseGame(game interactor.GameDTO) (output string) {
	output += "  "
	for j := range len(game.BoardCells) {
		output += fmt.Sprintf("%2s", shared.NumToExcelColumn(j))
	}
	output += "\n"
	for i := range game.BoardCells {
		output += fmt.Sprintf("%2d", i)
		for _, cell := range game.BoardCells[i] {
			if cell.IsFlagged {
				output += " x︎"
				continue
			}
			output += fmt.Sprintf("%2s", cellToStr(cell))
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
	return fmt.Sprint(cell.BombCount)
}
func (c *CLIController) debugGame(game interactor.GameDTO) (output string) {
	output += "  "
	for j := range len(game.BoardCells) {
		output += fmt.Sprintf("%2s", shared.NumToExcelColumn(j))
	}
	output += "\n"
	for i := range game.BoardCells {
		output += fmt.Sprintf("%2d", i)
		for _, cell := range game.BoardCells[i] {
			output += fmt.Sprintf("%2s", cellToDebugStr(cell))
		}
		output += "\n"
	}
	return output
}

func cellToDebugStr(cell interactor.CellDTO) string {
	if cell.IsBomb {
		return "B"
	}
	return fmt.Sprint(cell.BombCount)
}

func stateToStr(state byte) string {
	switch state {
	case domain.GameStatePlaying:
		return "Playing"
	case domain.GameStateCompleted:
		return "Completed"
	case domain.GameStateFailed:
		return "Failed"
	}
	return "?"
}

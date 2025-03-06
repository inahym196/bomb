package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb/internal/domain"
	"github.com/inahym196/bomb/internal/interactor"
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
			boardWidth, bombCount, err := c.parseStartArgs(words)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			result, err := c.gi.InitGame(interactor.InitGameParam{
				BoardWidth: boardWidth,
				BombCount:  bombCount,
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
			param := si.GetHintParam{Cells: result.GameDTO.BoardCells}
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
		case "exit", "quit", "q":
			fmt.Println("終了中...")
			return
		case "help", "h":
			fmt.Println(heredoc.Doc(`
			Available Commands:
			  > start <boardWidth> <bombCount>  Start Game
			  > show                            Show board
			  > open <row> <col>                Open cell
			  > help                            Show this help message
			  > exit                            Exit the program
			`))
		default:
			fmt.Printf("\"%s\"は無効なコマンドです. \"help\"コマンドを確認してください\n\n", text)
		}
	}
}

func (c *CLIController) parseStartArgs(words []string) (boardWidth, bombCount int, err error) {
	if len(words) != 3 {
		return 0, 0, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	boardWidth, err = strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, fmt.Errorf("boardWidthの値が不正です. 数字を入力してください")
	}
	bombCount, err = strconv.Atoi(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("bombCountの値が不正です. 数字を入力してください")
	}
	return boardWidth, bombCount, nil
}

func (c *CLIController) parseOpenArgs(words []string) (row, col int, err error) {
	if len(words) != 3 {
		return 0, 0, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	row, err = strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, fmt.Errorf("rowの値が不正です. 数字を入力してください")
	}
	col, err = strconv.Atoi(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("columnの値が不正です. 数字を入力してください")
	}
	return row, col, nil
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
	return fmt.Sprint(cell.BombCount)
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
	return fmt.Sprint(cell.BombCount)
}

func stateToStr(state byte) string {
	switch state {
	case domain.GameStateReady:
		return "Ready"
	case domain.GameStatePlaying:
		return "Playing"
	case domain.GameStateCompleted:
		return "Completed"
	case domain.GameStateFailed:
		return "Failed"
	}
	return "?"
}

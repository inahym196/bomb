package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/inahym196/bomb/pkg/shared"
)

const (
	GameModeUndefined byte = iota
	GameModeEasy
	GameModeNormal
	//GameModeHard
	GameModeCustom
)

func (c *CLIController) parseStartGameModeArgs(words []string) (mode byte, err error) {
	if len(words) < 2 {
		return GameModeUndefined, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	switch words[1] {
	case "easy":
		return GameModeEasy, nil
	case "normal":
		return GameModeNormal, nil
	// まだwidthとheightを個別に選択できないのでhard(30*16)は実装できない
	//case "hard":
	//	return GameModeHard, nil
	case "custom":
		return GameModeCustom, nil
	default:
		return GameModeUndefined, fmt.Errorf("start [easy, normal, custom] のいずれかを指定してください")
	}
}

func (c *CLIController) parseStartCustomModeArgs(words []string) (boardWidth, bombCount int, err error) {
	if words[1] != "custom" {
		log.Fatalln("この関数はcustomModeでのみ呼び出し可能です")
	}
	if len(words) != 4 {
		return 0, 0, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	boardWidth, err = strconv.Atoi(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("boardWidthの値が不正です. 数字を入力してください")
	}
	bombCount, err = strconv.Atoi(words[3])
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
	col, err = shared.ExcelColumnToNum(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("columnの値が不正です. アルファベット[a-zA-Z]を入力してください")
	}
	return row, col, nil
}

func (c *CLIController) parseCheckArgs(words []string) (row, col int, err error) {
	if len(words) != 3 {
		return 0, 0, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	row, err = strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, fmt.Errorf("rowの値が不正です. 数字を入力してください")
	}
	col, err = shared.ExcelColumnToNum(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("columnの値が不正です. アルファベット[a-zA-Z]を入力してください")
	}
	return row, col, nil
}

func (c *CLIController) parseUnCheckArgs(words []string) (row, col int, err error) {
	if len(words) != 3 {
		return 0, 0, fmt.Errorf("引数の数が不正です. \"help\"コマンドを確認してください")
	}
	row, err = strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, fmt.Errorf("rowの値が不正です. 数字を入力してください")
	}
	col, err = shared.ExcelColumnToNum(words[2])
	if err != nil {
		return 0, 0, fmt.Errorf("columnの値が不正です. アルファベット[a-zA-Z]を入力してください")
	}
	return row, col, nil
}

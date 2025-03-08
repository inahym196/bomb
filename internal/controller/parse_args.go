package controller

import (
	"fmt"
	"strconv"
)

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

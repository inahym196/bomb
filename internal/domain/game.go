package domain

import (
	"fmt"
	"math/rand"
	"reflect"
	"slices"
)

const (
	GameStateReady byte = iota
	GameStatePlaying
	GameStateFinished
)

const (
	GameResultUndefined byte = iota
	GameResultSuccess
	GameResultFailed
)

type Game struct {
	result     byte
	state      byte
	board      *Board
	boardWidth int
	bombCount  int
}

func (g *Game) GetBoard() *Board { return g.board }

func NewGame(boardWidth, bombCount int) (*Game, error) {
	if boardWidth < 2 {
		return nil, fmt.Errorf("boardWidthの数は2以上を指定してください")
	}
	maxBombCount := (boardWidth - 1) * (boardWidth - 1)
	if maxBombCount < bombCount {
		return nil, fmt.Errorf("boardWidthが%dの時、bombCountは%d以下を指定してください", boardWidth, maxBombCount)
	}
	return &Game{
		result:     GameResultUndefined,
		state:      GameStateReady,
		board:      NewBoard(boardWidth),
		boardWidth: boardWidth,
		bombCount:  bombCount,
	}, nil
}

type position struct {
	row int
	col int
}

func (g *Game) OpenCell(row, col int) error {
	if err := g.board.OpenCell(row, col); err != nil {
		return err
	}
	openedCell := g.board.GetCells()[row][col]
	if openedCell.IsBomb() {
		g.state = GameStateFinished
		g.result = GameResultFailed
		return nil
	}
	if g.state == GameStateReady {
		poss := NewRandomPositions(g.bombCount, g.boardWidth, position{row, col})
		for _, pos := range poss {
			g.board.SetBomb(pos.row, pos.col)
		}
		g.state = GameStatePlaying
	}
	return nil
}

func NewRandomPositions(n, maxN int, except position) []position {
	poss := make([]position, 0, n)
	for len(poss) != n {
		pos := position{rand.Intn(maxN), rand.Intn(maxN)}
		if !reflect.DeepEqual(pos, except) && !slices.Contains(poss, pos) {
			poss = append(poss, pos)
		}
	}
	return poss
}

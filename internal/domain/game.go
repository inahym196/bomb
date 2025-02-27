package domain

import (
	"math/rand"
	"reflect"
	"slices"
)

const (
	GameStateUninitialized byte = iota
	GameStateReady
	GameStatePlaying
	GameStateFinished
)

type Game struct {
	state      byte
	board      *Board
	boardWidth int
	bombCount  int
}

func (g *Game) GetBoard() *Board { return g.board }

func NewGame(boardWidth, bombCount int) *Game {
	return &Game{
		state:      GameStateUninitialized,
		board:      NewBoard(boardWidth),
		boardWidth: boardWidth,
		bombCount:  bombCount,
	}
}

type position struct {
	row int
	col int
}

func (g *Game) OpenCell(row, col int) error {
	if err := g.board.OpenCell(row, col); err != nil {
		return err
	}
	if g.state == GameStateUninitialized {
		poss := NewRandomPositions(g.bombCount, g.boardWidth, position{row, col})
		for _, pos := range poss {
			g.board.SetBomb(pos.row, pos.col)
		}
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

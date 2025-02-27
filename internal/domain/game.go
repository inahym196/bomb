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

type GameOption struct {
	BoardWidth int
	BombCount  int
}

func NewGame(opt *GameOption) *Game {
	if opt == nil {
		opt = &GameOption{
			BoardWidth: 9,
			BombCount:  10,
		}
	}
	return &Game{
		state:      GameStateUninitialized,
		board:      NewBoard(opt.BoardWidth),
		boardWidth: opt.BoardWidth,
		bombCount:  opt.BombCount,
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
		poss := NewRandomPositions(g.bombCount, g.boardWidth-1, position{row, col})
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

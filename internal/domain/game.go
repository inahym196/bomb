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
	GameStateCompleted
	GameStateFailed
)

type Game struct {
	state      byte
	board      *Board
	boardWidth int
	bombCount  int
}

func (g *Game) GetBoard() *Board { return g.board }
func (g *Game) GetState() byte   { return g.state }

func NewGame(boardWidth, bombCount int) (*Game, error) {
	if boardWidth < 2 {
		return nil, fmt.Errorf("boardWidthの数は2以上を指定してください")
	}
	maxBombCount := (boardWidth - 1) * (boardWidth - 1)
	if maxBombCount < bombCount {
		return nil, fmt.Errorf("boardWidthが%dの時、bombCountは%d以下を指定してください", boardWidth, maxBombCount)
	}
	return &Game{
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
	// TODO: state管理は別の関数に分けたい
	if g.bombCount == g.board.GetClosedCellCount() {
		g.state = GameStateCompleted
		return nil
	}
	openedCell := g.board.GetCells()[row][col]
	if openedCell.IsBomb() {
		g.state = GameStateFailed
		return nil
	}
	if g.state == GameStateReady {
		poss := g.newRandomPositions(position{row, col})
		for _, pos := range poss {
			g.board.SetBomb(pos.row, pos.col)
			g.incrementBombCountArroundBomb(pos.row, pos.col, g.board.IncrementBombCount)
		}
		g.state = GameStatePlaying
	}
	return nil
}

func (g *Game) newRandomPositions(except position) []position {
	n := g.bombCount
	maxN := g.boardWidth
	poss := make([]position, 0, n)
	for len(poss) != n {
		pos := position{rand.Intn(maxN), rand.Intn(maxN)}
		if !reflect.DeepEqual(pos, except) && !slices.Contains(poss, pos) {
			poss = append(poss, pos)
		}
	}
	return poss
}

func (g *Game) incrementBombCountArroundBomb(row, col int, incrementFunc func(row, col int) error) {
	cells := g.board.GetCells()
	dirs := []struct{ dy, dx int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for i := range dirs {
		nx, ny := col+dirs[i].dx, row+dirs[i].dy
		if g.board.inBoard(nx, ny) && !cells[ny][nx].IsBomb() {
			incrementFunc(ny, nx)
		}
	}
}

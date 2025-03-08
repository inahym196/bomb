package domain

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/inahym196/bomb/pkg/shared"
)

const (
	GameStateReady byte = iota
	GameStatePlaying
	GameStateCompleted
	GameStateFailed
)

type Game struct {
	state     byte
	board     *Board
	bombCount int
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
		state:     GameStateReady,
		board:     NewBoard(boardWidth),
		bombCount: bombCount,
	}, nil
}

func (g *Game) OpenCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	if g.state == GameStateReady {
		g.setRandomBombs(pos)
		g.state = GameStatePlaying
	}
	err := g.board.OpenCell(pos)
	if err != nil {
		return err
	}
	g.updateState(pos)
	return nil
}

func (g *Game) isFinished() bool {
	return g.state == GameStateCompleted || g.state == GameStateFailed
}

func (g *Game) setRandomBombs(except shared.Position) {
	poss := g.newRandomPositions(except)
	for _, pos := range poss {
		g.board.SetBomb(pos)
		g.incrementBombCountArroundBomb(pos, g.board.IncrementBombCount)
	}
}

func (g *Game) newRandomPositions(except shared.Position) []shared.Position {
	n := g.bombCount
	maxN := g.GetBoard().GetWidth()
	var poss []shared.Position
	for len(poss) != n {
		pos := shared.NewRandomPosition(maxN)
		if !reflect.DeepEqual(pos, except) && !slices.Contains(poss, pos) {
			poss = append(poss, pos)
		}
	}
	return poss
}

func (g *Game) incrementBombCountArroundBomb(pos shared.Position, incrementFunc func(pos shared.Position) error) {
	cells := g.board.GetCells()
	pos.ForEachNeighbor(func(p shared.Position) {
		if g.board.inBoard(p) && !cells[p.Y][p.X].IsBomb() {
			incrementFunc(p)
		}
	})
}

func (g *Game) updateState(pos shared.Position) {
	switch {
	case g.bombCount == g.board.GetClosedCellCount():
		g.state = GameStateCompleted
	case g.board.MustGetCellAt(pos).IsBomb():
		g.state = GameStateFailed
	}
}

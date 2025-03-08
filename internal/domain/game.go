package domain

import (
	"fmt"

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
		bombPositions := shared.NewUniqueRandomPositionsWithout(g.bombCount, g.board.GetWidth(), pos)
		g.board.SetBombs(bombPositions)
		g.state = GameStatePlaying
	}
	err := g.board.OpenCell(pos)
	if err != nil {
		return err
	}
	if g.board.MustGetCellAt(pos).IsBomb() {
		g.state = GameStateFailed
	}
	if g.bombCount == g.board.GetClosedCellCount() {
		g.state = GameStateCompleted
	}
	return nil
}

func (g *Game) isFinished() bool {
	return g.state == GameStateCompleted || g.state == GameStateFailed
}

func (g *Game) CheckCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	return g.board.CheckCell(pos)
}

func (g *Game) UnCheckCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	err := g.board.UnCheckCell(pos)
	if err != nil {
		return err
	}
	return nil
}

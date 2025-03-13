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
	state      byte
	bombField  *BombField
	bombCount  int
	boardWidth int
}

func (g *Game) GetBombField() *BombField { return g.bombField }
func (g *Game) GetState() byte           { return g.state }

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
		bombField:  NewBombField(boardWidth),
		bombCount:  bombCount,
		boardWidth: boardWidth,
	}, nil
}

func (g *Game) OpenCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	if g.state == GameStateReady {
		bombPositions := shared.NewUniqueRandomPositionsWithout(g.bombCount, g.boardWidth, pos)
		g.bombField.SetBombs(bombPositions)
		g.state = GameStatePlaying
	}
	bursted, err := g.bombField.OpenCell(pos)
	if err != nil {
		return err
	}
	if bursted {
		g.state = GameStateFailed
	}
	if g.bombCount == g.bombField.GetClosedCellCount() {
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
	return g.bombField.CheckCell(pos)
}

func (g *Game) UnCheckCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	err := g.bombField.UnCheckCell(pos)
	if err != nil {
		return err
	}
	return nil
}

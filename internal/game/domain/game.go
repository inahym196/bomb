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
	totalBomb  int
	boardWidth int
}

func (g *Game) GetBombField() *BombField { return g.bombField }
func (g *Game) GetState() byte           { return g.state }

func NewGame(boardWidth, totalBomb int) (*Game, error) {
	bombField, err := NewBombField(boardWidth, totalBomb)
	if err != nil {
		return nil, err
	}
	return &Game{
		state:      GameStateReady,
		bombField:  bombField,
		totalBomb:  totalBomb,
		boardWidth: boardWidth,
	}, nil
}

func (g *Game) OpenCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	if g.state == GameStateReady {
		bombPositions := shared.NewUniqueRandomPositionsWithout(g.totalBomb, g.boardWidth, pos)
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
	if g.bombField.IsPeaceFul() {
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

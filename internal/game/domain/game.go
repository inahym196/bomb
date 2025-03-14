package domain

import (
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

const (
	GameStatePlaying byte = iota
	GameStateCompleted
	GameStateFailed
)

type Game struct {
	state     byte
	bombField *BombField
}

func (g *Game) GetBombField() *BombField { return g.bombField }
func (g *Game) GetState() byte           { return g.state }
func (g *Game) isFinished() bool         { return g.state == GameStateCompleted || g.state == GameStateFailed }

func NewGame(boardWidth, totalBomb int) (*Game, error) {
	bombField, err := NewBombField(boardWidth, totalBomb)
	if err != nil {
		return nil, err
	}
	return &Game{
		state:     GameStatePlaying,
		bombField: bombField,
	}, nil
}

func (g *Game) OpenCell(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
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

func (g *Game) Flag(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	return g.bombField.Flag(pos)
}

func (g *Game) UnFlag(pos shared.Position) error {
	if g.isFinished() {
		return fmt.Errorf("ゲームはすでに終了しています")
	}
	return g.bombField.UnFlag(pos)
}

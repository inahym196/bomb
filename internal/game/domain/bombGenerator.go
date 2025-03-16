package domain

import "github.com/inahym196/bomb/pkg/shared"

type BombGenerator interface {
	GenerateWithout(pos shared.Position) map[shared.Position]struct{}
}

type defaultBombGenerator struct {
	totalBomb int
	width     int
}

func newDefaultBombGenerator(totalBomb, width int) *defaultBombGenerator {
	return &defaultBombGenerator{totalBomb, width}
}

func (bg *defaultBombGenerator) GenerateWithout(pos shared.Position) map[shared.Position]struct{} {
	return shared.NewUniqueRandomPositionsWithout(bg.totalBomb, bg.width, pos)
}

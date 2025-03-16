package domain

import "github.com/inahym196/bomb/pkg/shared"

type BombGenerator interface {
	GenerateWithout(pos shared.Position) map[shared.Position]struct{}
}

type bombGenerator struct {
	totalBomb int
	width     int
}

func newBombGenerator(totalBomb, width int) *bombGenerator { return &bombGenerator{totalBomb, width} }

func (bg *bombGenerator) GenerateWithout(pos shared.Position) map[shared.Position]struct{} {
	return shared.NewUniqueRandomPositionsWithout(bg.totalBomb, bg.width, pos)
}

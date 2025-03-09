package domain

import (
	"testing"

	"github.com/inahym196/bomb/pkg/shared"
)

func BenchmarkSetBomb(b *testing.B) {
	b.ResetTimer()
	width := 10
	bombCount := 9
	except := shared.NewRandomPosition(width)
	bombPositions := shared.NewUniqueRandomPositionsWithout(width, bombCount, except)
	board := NewBoard(width)
	for range b.N {
		board.SetBombs(bombPositions)
	}
}

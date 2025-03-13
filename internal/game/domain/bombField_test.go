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
	bombField := NewBombField(width)
	for range b.N {
		bombField.SetBombs(bombPositions)
	}
}

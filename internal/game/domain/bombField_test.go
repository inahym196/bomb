package domain

import (
	"testing"

	"github.com/inahym196/bomb/pkg/shared"
)

func BenchmarkSetBomb(b *testing.B) {
	b.ResetTimer()
	width := 10
	totalBomb := 9
	except := shared.NewRandomPosition(width)
	bombPositions := shared.NewUniqueRandomPositionsWithout(width, totalBomb, except)
	bombField, _ := NewBombField(width, totalBomb)
	for range b.N {
		bombField.SetBombs(bombPositions)
	}
}

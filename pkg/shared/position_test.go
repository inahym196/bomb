package shared

import (
	"testing"
)

const n = 30

func BenchmarkNewUniqueRandomPositionsWithout(b *testing.B) {
	for range b.N {
		NewUniqueRandomPositionsWithout(n, n, NewPosition(0, 0))
	}
}

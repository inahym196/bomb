package shared

import (
	"slices"
	"testing"
)

const n = 30

func BenchmarkNewUniqueRandomPositionsWithout(b *testing.B) {
	for range b.N {
		NewUniqueRandomPositionsWithout(n, n, NewPosition(0, 0))
	}
}

func TestNewUniqueRandomPositionsWithout(t *testing.T) {
	type args struct {
		n      int
		maxN   int
		except Position
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "n=9", args: args{n: 9, maxN: 9, except: NewPosition(0, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUniqueRandomPositionsWithout(tt.args.n, tt.args.maxN, tt.args.except)
			if len(got) != tt.args.n {
				t.Errorf("len = %d, want %d", len(got), tt.args.n)
			}
			if slices.Contains(got, tt.args.except) {
				t.Errorf("It contains except position %v", tt.args.except)
			}
			for i, pos := range got {
				if slices.Contains(got[i+1:], pos) {
					t.Errorf("It contains same position %v", got)
				}
			}
		})
	}
}

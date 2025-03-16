package domain

import (
	"reflect"
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
		bombField.setBombs(bombPositions)
	}
}

type testBombGenerator struct{}

var _ BombGenerator = testBombGenerator{}

func (tbg testBombGenerator) GenerateWithout(pos shared.Position) map[shared.Position]struct{} {
	return nil
}

func TestNewBombField(t *testing.T) {
	type args struct {
		width     int
		totalBomb int
	}
	tests := []struct {
		name    string
		args    args
		want    *BombField
		wantErr bool
	}{
		{
			name: "fail case: width < 2",
			args: args{
				width:     1,
				totalBomb: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail case: invalid totalbomb",
			args: args{
				width:     2,
				totalBomb: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success case",
			args: args{
				width:     2,
				totalBomb: 1,
			},
			want:    &BombField{NewBoard(2), newBombCounter(2), newBombGenerator(1, 2), newFieldState(1, 2)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBombField(tt.args.width, tt.args.totalBomb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBombField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBombField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBombField_WithBombGenerator(t *testing.T) {
	type args struct {
		width     int
		totalBomb int
		bg        BombGenerator
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				width:     10,
				totalBomb: 9,
				bg:        testBombGenerator{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf, err := NewBombField(tt.args.width, tt.args.totalBomb)
			if err != nil {
				t.Error(err)
			}
			if reflect.TypeOf(bf.bombGenerator) != reflect.TypeOf(&bombGenerator{}) {
				t.Fatalf("got %T, expect %T", bf.bombGenerator, bombGenerator{})
			}
			bf.WithBombGenerator(tt.args.bg)
			if reflect.TypeOf(bf.bombGenerator) != reflect.TypeOf(tt.args.bg) {
				t.Errorf("got %T, expect %T", bf.bombGenerator, tt.args.bg)
			}
		})
	}
}

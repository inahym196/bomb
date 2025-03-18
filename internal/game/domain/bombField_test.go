package domain

import (
	"reflect"
	"testing"

	"github.com/inahym196/bomb/pkg/shared"
)

func BenchmarkSetBomb(b *testing.B) {
	width := 10
	totalBomb := 9
	except := shared.NewRandomPosition(width)
	bombPositions := shared.NewUniqueRandomPositionsWithout(width, totalBomb, except)
	bombField, _ := NewBombField(width, totalBomb)
	b.ResetTimer()
	for range b.N {
		bombField.setBombs(bombPositions)
	}
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
			want:    &BombField{NewBoard(2), newBombCounter(2), newDefaultBombGenerator(1, 2), newFieldState(1, 2)},
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

type testBombGenerator struct{}

var _ BombGenerator = testBombGenerator{}

func (tbg testBombGenerator) GenerateWithout(pos shared.Position) []shared.Position {
	return []shared.Position{shared.NewPosition(0, 0)}
}

type testBombGenerator2 struct{}

var _ BombGenerator = testBombGenerator2{}

func (tbg testBombGenerator2) GenerateWithout(pos shared.Position) []shared.Position {
	return []shared.Position{shared.NewPosition(1, 0)}
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
			if reflect.TypeOf(bf.bombGenerator) != reflect.TypeOf(&defaultBombGenerator{}) {
				t.Fatalf("got %T, expect %T", bf.bombGenerator, &defaultBombGenerator{})
			}
			bf.WithBombGenerator(tt.args.bg)
			if reflect.TypeOf(bf.bombGenerator) != reflect.TypeOf(tt.args.bg) {
				t.Errorf("got %T, expect %T", bf.bombGenerator, tt.args.bg)
			}
		})
	}
}

func TestBombField_OpenCell(t *testing.T) {
	t.Run("board外を選択するとエラー", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(2, 0),
		}
		bf := &BombField{
			board:         NewBoard(2),
			bombCounter:   newBombCounter(2),
			bombGenerator: testBombGenerator{},
			state:         newFieldState(1, 2),
		}
		wantErr := true

		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
	})
	t.Run("初回SafeCell(size=2)", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(1, 1),
		}
		bf := &BombField{
			board:         NewBoard(2),
			bombCounter:   newBombCounter(2),
			bombGenerator: testBombGenerator{},
			state:         newFieldState(1, 2),
		}
		wantErr := false
		wantCounts := [][]int{{-1, 1}, {1, 1}}

		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(bf.bombCounter.counts, wantCounts) {
			t.Errorf("counts: got %v, want %v", bf.bombCounter.counts, wantCounts)
		}
	})
	t.Run("初回SafeCell(size=3)", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(1, 1),
		}
		bf := &BombField{
			board:         NewBoard(3),
			bombCounter:   newBombCounter(3),
			bombGenerator: testBombGenerator{},
			state:         newFieldState(1, 3),
		}
		wantErr := false
		wantCounts := [][]int{{-1, 1, 0}, {1, 1, 0}, {0, 0, 0}}

		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
		if !reflect.DeepEqual(bf.bombCounter.counts, wantCounts) {
			t.Errorf("counts: got %v, want %v", bf.bombCounter.counts, wantCounts)
		}
	})
	t.Run("BombCount=0を開くと周りも開放される", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(2, 2),
		}
		bf := &BombField{
			board:         NewBoard(3),
			bombCounter:   newBombCounter(3),
			bombGenerator: testBombGenerator{},
			state:         newFieldState(1, 3),
		}
		wantErr := false
		opened, _ := NewSafeCell().Open()
		closed := NewBombCell()
		cells := [][]Cell{{closed, opened, opened}, {opened, opened, opened}, {opened, opened, opened}}
		wantBoard := &board{width: 3, cells: cells}
		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
		if !bf.board.Equals(wantBoard) {
			t.Errorf("board: got %v, want %v", bf.board, wantBoard)
		}
	})
	t.Run("BombCount=0を開くと周りも開放される 2", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(2, 2),
		}
		bf := &BombField{
			board:         NewBoard(3),
			bombCounter:   newBombCounter(3),
			bombGenerator: testBombGenerator2{},
			state:         newFieldState(1, 3),
		}
		wantErr := false
		closedSafe := NewSafeCell()
		openedSafe, _ := closedSafe.Open()
		closedBomb := NewBombCell()
		cells := [][]Cell{
			{closedSafe, closedBomb, closedSafe},
			{openedSafe, openedSafe, openedSafe},
			{openedSafe, openedSafe, openedSafe},
		}
		wantBoard := &board{width: 3, cells: cells}
		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
		if !bf.board.Equals(wantBoard) {
			t.Errorf("board: \ngot %v,\nwant %v", bf.board, wantBoard)
		}
	})
	t.Run("BombCellも開ける", func(t *testing.T) {
		args := struct{ pos shared.Position }{
			pos: shared.NewPosition(0, 0),
		}
		bf := &BombField{
			board:         NewBoard(2),
			bombCounter:   newBombCounter(2),
			bombGenerator: testBombGenerator{},
			state:         newFieldState(1, 2),
		}
		wantErr := false

		if err := bf.OpenCell(args.pos); (err != nil) != wantErr {
			t.Errorf("BombField.OpenCell() error = %v, wantErr %v", err, wantErr)
		}
		if !bf.state.IsBursted() {
			t.Fatal("Arrange失敗: 開いたのはBombCellじゃない")
		}
	})
}

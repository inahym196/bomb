package domain

const (
	CellUndefined byte = iota
	CellClosed
	CellBomb
	CellOpen
)

type Cell struct {
	isOpened bool
	isBomb   bool
}

func (c Cell) IsOpened() bool { return c.isOpened }
func (c Cell) IsBomb() bool   { return c.isBomb }

func NewCell(isOpened, isBomb bool) Cell {
	return Cell{isOpened, isBomb}
}

type Board struct {
	cells [][]Cell
}

func (b *Board) GetCells() [][]Cell { return b.cells }

func NewBoard(width int) *Board {
	board := &Board{}
	board.cells = initCells(width)
	return board
}

func initCells(width int) [][]Cell {
	cells := make([][]Cell, width)
	for i := range width {
		cells[i] = make([]Cell, width)
		for j := range width {
			cells[i][j] = NewCell(false, false)
		}
	}
	return cells
}

type Position struct {
	Row int
	Col int
}

func (b *Board) SetBombs(positionList []Position) {
	for _, pos := range positionList {
		b.cells[pos.Row][pos.Col] = NewCell(false, true)
	}
}

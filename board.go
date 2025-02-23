package bomb

const (
	CellUndefined byte = iota
	CellClosed
	CellBomb
	CellOpen
)

type Cell struct {
	state byte
}

func NewCell(state byte) Cell {
	if state != CellClosed && state != CellBomb && state != CellOpen {
		return Cell{CellUndefined}
	}
	return Cell{state}
}

type Board struct {
	cells [][]Cell
}

func NewBoard(width int) *Board {
	board := &Board{}
	board.init(width)
	return board
}

func (b *Board) init(width int) {
	cells := make([][]Cell, width)
	for i, row := range cells {
		cells[i] = make([]Cell, width)
		for j := range row {
			cells[i][j] = NewCell(CellClosed)
		}
	}
	b.cells = cells
}

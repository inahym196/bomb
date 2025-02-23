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
	board.cells = initCells(width)
	return board
}

func initCells(width int) [][]Cell {
	cells := make([][]Cell, width)
	for i, row := range cells {
		cells[i] = make([]Cell, width)
		for j := range row {
			cells[i][j] = NewCell(CellClosed)
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
		b.cells[pos.Row][pos.Col] = NewCell(CellBomb)
	}
}

package domain

const (
	FieldStateAllClosed = iota
	FieldStateNotBurst
	FieldStateBursted
	FieldStatePeaceful
)

type fieldState struct {
	totalBomb       int
	width           int
	closedCellCount int
	state           byte
}

func newFieldState(totalBomb, width int) *fieldState {
	return &fieldState{totalBomb: totalBomb, width: width, closedCellCount: width * width, state: FieldStateAllClosed}
}

func (fs *fieldState) IsPeaceFul() bool  { return fs.state == FieldStatePeaceful }
func (fs *fieldState) IsAllClosed() bool { return fs.state == FieldStateAllClosed }
func (fs *fieldState) IsBursted() bool   { return fs.state == FieldStateBursted }
func (fs *fieldState) Burst()            { fs.state = FieldStateBursted }

func (fs *fieldState) DecrementClosedCell() {
	fs.closedCellCount--
	if fs.state == FieldStateAllClosed {
		fs.state = FieldStateNotBurst
		return
	}
	if fs.closedCellCount == fs.totalBomb {
		fs.state = FieldStatePeaceful
		return
	}
}

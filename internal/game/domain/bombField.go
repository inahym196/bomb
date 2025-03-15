package domain

import (
	"container/list"
	"fmt"

	"github.com/inahym196/bomb/pkg/errorutil"
	"github.com/inahym196/bomb/pkg/shared"
)

type BombField struct {
	board         *board
	bombCounter   *bombCounter
	bombGenerator *bombGenerator
	state         *fieldState
}

func (bf *BombField) GetCells() [][]Cell     { return bf.board.GetCells() }
func (bf *BombField) GetBombCounts() [][]int { return bf.bombCounter.GetBombCounts() }

func (bf *BombField) IsPeaceFul() bool { return bf.state.IsPeaceFul() }
func (bf *BombField) IsBursted() bool  { return bf.state.IsBursted() }

func NewBombField(width int, totalBomb int) (*BombField, error) {
	if width < 2 {
		return nil, fmt.Errorf("widthは2以上を指定してください")
	}
	maxBombLimit := (width - 1) * (width - 1)
	if maxBombLimit < totalBomb {
		return nil, fmt.Errorf("widthが%dの時、totalBombは%d以下を指定してください", width, maxBombLimit)
	}
	return &BombField{
		board:         NewBoard(width),
		bombCounter:   newBombCounter(width),
		bombGenerator: newBombGenerator(totalBomb, width),
		state:         newFieldState(totalBomb, width),
	}, nil
}

func (bf *BombField) OpenCell(pos shared.Position) (err error) {
	if bf.state.IsAllClosed() {
		bf.setBombs(bf.bombGenerator.GenerateWithout(pos))
	}
	err = bf.openCell(pos)
	if err != nil {
		return err
	}
	if !bf.state.IsBursted() {
		bf.expandSafeArea(pos)
	}
	return nil
}

func (bf *BombField) setBombs(positions map[shared.Position]struct{}) error {
	for pos := range positions {
		if !bf.board.contains(pos) {
			return fmt.Errorf("ボード外のポジションは指定できません")
		}
		bf.board.setCellAt(pos, NewBombCell())
		bf.bombCounter.SetBombCount(pos)
	}
	return nil
}

func (bf *BombField) openCell(pos shared.Position) error {
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return err
	}
	openedCell, err := cell.Open()
	if err != nil {
		return err
	}
	bf.board.setCellAt(pos, openedCell)
	bf.state.DecrementClosedCell()
	if openedCell.IsBomb() {
		bf.state.Burst()
	}
	return nil
}

func (bf *BombField) expandSafeArea(pos shared.Position) {
	visited := make([][]bool, bf.board.GetWidth())
	for i := range visited {
		visited[i] = make([]bool, bf.board.GetWidth())
	}
	queue := list.New()
	queue.PushBack(pos)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		pos := front.Value.(shared.Position)
		if visited[pos.Y][pos.X] {
			continue
		}
		visited[pos.Y][pos.X] = true
		_ = bf.openCell(pos)
		if bf.bombCounter.IsNeighborsSafe(pos) {
			pos.ForEachNeighbor(func(p shared.Position) {
				cell, err := bf.board.GetCellAt(p)
				if err == nil && !visited[p.Y][p.X] && cell.CanOpen() {
					queue.PushBack(p)
				}
			})
		}
	}
}

func (bf *BombField) Flag(pos shared.Position) (err error) {
	defer errorutil.Wrap(&err, "flag(%v)", pos)
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return err
	}
	flagged, err := cell.Flagged()
	if err != nil {
		return err
	}
	bf.board.setCellAt(pos, flagged)
	return nil
}

func (bf *BombField) UnFlag(pos shared.Position) (err error) {
	defer errorutil.Wrap(&err, "unflag(%v)", pos)
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return err
	}
	unflagged, err := cell.UnFlagged()
	if err != nil {
		return err
	}
	bf.board.setCellAt(pos, unflagged)
	return nil
}

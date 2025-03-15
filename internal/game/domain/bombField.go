package domain

import (
	"container/list"
	"fmt"

	"github.com/inahym196/bomb/pkg/errorutil"
	"github.com/inahym196/bomb/pkg/shared"
)

type bombGenerator struct {
	totalBomb int
	width     int
}

func newBombGenerator(totalBomb, width int) *bombGenerator { return &bombGenerator{totalBomb, width} }

func (bg *bombGenerator) GenerateWithout(pos shared.Position) map[shared.Position]struct{} {
	return shared.NewUniqueRandomPositionsWithout(bg.totalBomb, bg.width, pos)
}

type BombField struct {
	board           *board
	closedCellCount int
	bombCounts      [][]int
	totalBomb       int
	bombGenerator   *bombGenerator
}

func (bf *BombField) GetCells() [][]Cell     { return bf.board.GetCells() }
func (bf *BombField) GetBombCounts() [][]int { return bf.bombCounts }

func (bf *BombField) IsPeaceFul() bool                  { return bf.closedCellCount == bf.totalBomb }
func (bf *BombField) isAllClosed() bool                 { return bf.closedCellCount == bf.board.width*bf.board.width }
func (bf *BombField) Contains(pos shared.Position) bool { return bf.board.contains(pos) }

func NewBombField(width int, totalBomb int) (*BombField, error) {
	if width < 2 {
		return nil, fmt.Errorf("widthは2以上を指定してください")
	}
	maxBombLimit := (width - 1) * (width - 1)
	if maxBombLimit < totalBomb {
		return nil, fmt.Errorf("widthが%dの時、totalBombは%d以下を指定してください", width, maxBombLimit)
	}
	return &BombField{
		board:           NewBoard(width),
		closedCellCount: width * width,
		bombCounts:      initBombCounts(width),
		totalBomb:       totalBomb,
		bombGenerator:   newBombGenerator(totalBomb, width),
	}, nil
}

func initBombCounts(width int) [][]int {
	totalBombs := make([][]int, width)
	for i := range width {
		totalBombs[i] = make([]int, width)
		for j := range width {
			totalBombs[i][j] = 0
		}
	}
	return totalBombs
}

func (bf *BombField) OpenCell(pos shared.Position) (bursted bool, err error) {
	if bf.isAllClosed() {
		bf.setBombs(bf.bombGenerator.GenerateWithout(pos))
	}
	bursted, err = bf.openCell(pos)
	if err != nil {
		return false, err
	}
	if bursted {
		return true, nil
	}
	bf.expandSafeArea(pos)
	return false, nil
}

func (bf *BombField) setBombs(positions map[shared.Position]struct{}) error {
	for pos := range positions {
		if !bf.board.contains(pos) {
			return fmt.Errorf("ボード外のポジションは指定できません")
		}
		bf.board.setCellAt(pos, NewBombCell())
		bf.bombCounts[pos.Y][pos.X] = -1
		bf.incrementBombCountForEachNeighbor(pos)
	}
	return nil
}

func (bf *BombField) incrementBombCountForEachNeighbor(pos shared.Position) {
	pos.ForEachNeighbor(func(p shared.Position) {
		if bf.board.contains(p) && bf.bombCounts[p.Y][p.X] != -1 {
			bf.bombCounts[p.Y][p.X]++
		}
	})
}

func (bf *BombField) openCell(pos shared.Position) (bursted bool, err error) {
	cell, err := bf.board.GetCellAt(pos)
	if err != nil {
		return false, err
	}
	openedCell, err := cell.Open()
	if err != nil {
		return false, err
	}
	bf.board.setCellAt(pos, openedCell)
	bf.closedCellCount--
	return openedCell.IsBomb(), nil
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
		_, _ = bf.openCell(pos)
		if bf.bombCounts[pos.Y][pos.X] == 0 {
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

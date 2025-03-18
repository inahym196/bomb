package domain

import (
	"fmt"

	"github.com/inahym196/bomb/pkg/shared"
)

type solver struct {
	cells    map[shared.Position]OpenCell
	theorems []Theorem
}

func NewSolver(cells map[shared.Position]OpenCell) solver {
	return solver{
		cells: cells,
		theorems: []Theorem{
			theorem1{},
		},
	}
}

func (s solver) Solve() string {
	var str string
	for _, theorem := range s.theorems {
		solution := theorem.Apply(s.cells)
		if len(solution.Positions) == 0 {
			str += fmt.Sprintf("[%s]: 該当なし\n", s.theorems[0].GetDescription())
			continue
		}
		var result string
		switch solution.Result {
		case SolutionResultIsBomb:
			result = "Bombである"
		case SolutionResultIsNotBomb:
			result = "開放できる"
		}
		str += fmt.Sprintf("[%s]に従い、\n以下は全て%s\n", theorem.GetDescription(), result)
		for _, pos := range solution.Positions {
			str += fmt.Sprintf("%v\n", pos)
		}
		str += fmt.Sprint()
	}
	return str
}

const (
	SolutionResultIsNotBomb byte = iota
	SolutionResultIsBomb    byte = iota
)

type Solution struct {
	Positions []shared.Position
	Result    byte
}

type Theorem interface {
	GetDescription() string
	Apply(map[shared.Position]OpenCell) Solution
}

type theorem1 struct{}

func (t theorem1) GetDescription() string {
	return "shadyCellsがtotalBomb以下なら全部bomb"
}

func (t theorem1) Apply(cells map[shared.Position]OpenCell) Solution {
	poss := make([]shared.Position, 0, len(cells)/2)
	for _, opencell := range cells {
		shadyCells := opencell.GetShadyCellKeys()
		if len(shadyCells) <= opencell.GetBombCount() {
			poss = append(poss, shadyCells...)
		}
	}
	return Solution{Result: SolutionResultIsBomb, Positions: poss}
}

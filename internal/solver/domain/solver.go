package domain

import (
	"fmt"

	"github.com/inahym196/bomb/internal/game/interactor"
	"github.com/inahym196/bomb/pkg/shared"
)

type solver struct {
	cells    [][]interactor.CellDTO
	theorems []Theorem
}

func NewSolver(cells [][]interactor.CellDTO) solver {
	return solver{
		cells: cells,
		theorems: []Theorem{
			theorem1{},
			theorem2{},
		},
	}
}

func (s solver) Solve() string {
	var str string
	for _, theorem := range s.theorems {
		solution := theorem.Apply(s.cells)
		if len(solution.Positions) == 0 {
			str += fmt.Sprintf("[%s]: 該当なし\n", theorem.GetDescription())
			continue
		}
		var result string
		switch solution.Result {
		case SolutionResultIsBomb:
			result = "Bomb"
		case SolutionResultIsSafe:
			result = "安全"
		}
		str += fmt.Sprintf("[%s]に従い、\n以下は全て%sである\n", theorem.GetDescription(), result)
		for _, pos := range solution.Positions {
			str += fmt.Sprintf("(%d,%s)\n", pos.Y, shared.NumToExcelColumn(pos.X))
		}
		str += "\n"
	}
	return str
}

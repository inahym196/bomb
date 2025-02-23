package main

import (
	"github.com/inahym196/bomb/internal/controller"
	"github.com/inahym196/bomb/internal/domain"
)

func main() {

	opt := &domain.GameOption{
		BoardWidth: 8,
		Bombs:      []domain.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	}
	game := domain.NewGame(opt)
	controller.NewCLIController(game).Run()
}

package main

import (
	"github.com/inahym196/bomb"
	"github.com/inahym196/bomb/controller"
)

func main() {

	opt := &bomb.GameOption{
		BoardWidth: 8,
		Bombs:      []bomb.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	}
	game := bomb.NewGame(opt)
	controller.NewCLIController(game).Run()
}

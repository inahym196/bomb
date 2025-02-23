package main

import (
	"fmt"

	"github.com/inahym196/bomb"
)

func main() {

	opt := &bomb.GameOption{
		BoardWidth: 8,
		Bombs:      []bomb.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	}
	game := bomb.NewGame(opt)
	fmt.Printf("%v", game.GetBoard())
}

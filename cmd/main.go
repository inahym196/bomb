package main

import (
	"fmt"

	"github.com/inahym196/bomb"
)

func main() {
	opt := &bomb.GameOption{BoardWidth: 8}
	game := bomb.NewGame(opt)
	fmt.Printf("%v", game.GetBoard())
}

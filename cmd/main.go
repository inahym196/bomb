package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/inahym196/bomb"
	"github.com/inahym196/bomb/controller"
)

func main() {

	opt := &bomb.GameOption{
		BoardWidth: 8,
		Bombs:      []bomb.Position{{Row: 0, Col: 0}, {Row: 1, Col: 2}},
	}
	game := bomb.NewGame(opt)
	controller := controller.NewCLIController(game)
	run(controller)
}

func run(controller *controller.CLIController) {

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("---")
		fmt.Print("Enter command> ")
		input.Scan()
		switch input.Text() {
		case "show", "s":
			fmt.Print(controller.GetBoard())
		case "exit", "quit":
			return
		case "help", "h":
			fmt.Print(heredoc.Doc(`
			Usage:
			  > help              this message
			  > show              show board
			  > open <row> <col>  open cell
			  > exit              end game
			`))
		default:
			fmt.Printf("\"%s\" is invalid command. Use \"help\"\n", input.Text())
		}
	}
}

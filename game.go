package bomb

type Game struct {
	board *Board
}

type GameOption struct {
	BoardWidth int
	Bombs      []Position
}

func NewGame(opt *GameOption) *Game {
	if opt == nil {
		opt = &GameOption{}
	}
	game := &Game{NewBoard(opt.BoardWidth)}
	game.board.SetBombs(opt.Bombs)
	return game
}

func (g *Game) GetBoard() *Board {
	return g.board
}

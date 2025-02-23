package bomb

type Game struct {
	board *Board
}

type GameOption struct {
	BoardWidth int
}

func NewGame(opt *GameOption) *Game {
	if opt == nil {
		opt = &GameOption{}
	}
	return &Game{
		board: NewBoard(opt.BoardWidth),
	}
}

func (g *Game) GetBoard() *Board {
	return g.board
}

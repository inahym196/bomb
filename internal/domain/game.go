package domain

type Game struct {
	board *Board
}

func (g *Game) GetBoard() *Board { return g.board }

type GameOption struct {
	BoardWidth int
	BombCount  int
}

func NewGame(opt *GameOption) *Game {
	if opt == nil {
		opt = &GameOption{}
	}
	game := &Game{NewBoard(opt.BoardWidth)}
	return game
}

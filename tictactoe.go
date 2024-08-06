package tictactoe

const (
	PlayerX = iota
	PlayerO
)

type Game struct {
	players []int
}

func NewGame() Game {
	return Game{
		players: []int{0, 0},
	}
}

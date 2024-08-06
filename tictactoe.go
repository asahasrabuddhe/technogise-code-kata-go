package tictactoe

import "errors"

const (
	PlayerX = iota
	PlayerO
)

var (
	ErrGameOver            = errors.New("game over")
	ErrFieldOccupied error = errors.New("field occupied")
)

type Game struct {
	players    []int
	turn       int
	isGameOver bool
}

func NewGame() Game {
	return Game{
		players: []int{0, 0},
		turn:    1,
	}
}

func (g *Game) TakeField(position int) error {
	if g.isGameOver {
		return ErrGameOver
	}

	move := 1 << position
	if g.players[PlayerX]&move != 0 || g.players[PlayerO]&move != 0 {
		// The position is already taken
		return ErrFieldOccupied
	}

	if g.turn == 1 {
		g.players[PlayerX] |= move
	} else {
		g.players[PlayerO] |= move
	}

	g.turn *= -1

	return nil
}

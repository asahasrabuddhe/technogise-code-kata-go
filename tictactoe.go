package tictactoe

import "errors"

const (
	PlayerX = iota
	PlayerO
)

// Win Conditions
var winConditions = []int{
	// diagonals
	0b100010001, // top-left
	0b001010100, // top-right
	// rows
	0b111000000, // top
	0b000111000, // middle
	0b000000111, // bottom
	// columns
	0b100100100, // left
	0b010010010, // middle
	0b001001001, // right
}

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

	g.checkGameResult()

	return nil
}

func (g *Game) checkGameResult() {
	if (g.players[PlayerX] | g.players[PlayerO]) == 0b111111111 {
		g.isGameOver = true
		return
	}

	for _, condition := range winConditions {
		// Player X wins
		if g.players[PlayerX]&condition == condition {
			g.isGameOver = true
			return
		}
		// Player O wins
		if g.players[PlayerO]&condition == condition {
			g.isGameOver = true
			return
		}
	}
}

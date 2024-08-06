package tictactoe

import (
	"errors"
	"testing"
)

func TestPlayers(t *testing.T) {
	t.Run("there are two players in the game, X and O", func(t *testing.T) {
		g := NewGame()
		if len(g.players) != 2 {
			t.Errorf("got %d g.players, want 2", len(g.players))
		}
	})

	t.Run(
		"consecutive turns alternate between the two players, X and O, until the game is over",
		func(t *testing.T) {
			g := NewGame()
			if g.turn != 1 {
				t.Errorf("g.turn = %d, want 1", g.turn)
			}

			_ = g.TakeField(0)
			if g.turn != -1 {
				t.Errorf("got player %v, want -1", g.turn)
			}

			_ = g.TakeField(1)
			if g.turn != 1 {
				t.Errorf("got player %v, want 1", g.turn)
			}

			_ = g.TakeField(2)
			if g.turn != -1 {
				t.Errorf("got player %v, want -1", g.turn)
			}

			_ = g.TakeField(3)
			if g.turn != 1 {
				t.Errorf("got player %v, want 1", g.turn)
			}

			g.isGameOver = true

			err := g.TakeField(4)
			if !errors.Is(err, ErrGameOver) {
				t.Errorf("got %v, want nil", err)
			}
		},
	)
}

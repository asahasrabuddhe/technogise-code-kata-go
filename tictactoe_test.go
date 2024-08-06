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

	t.Run("a player can take a field if not already taken", func(t *testing.T) {
		g := NewGame()

		err := g.TakeField(1)
		if err != nil {
			t.Errorf("got %v, want nil", err)
		}

		err = g.TakeField(1)
		if !errors.Is(err, ErrFieldOccupied) {
			t.Errorf("got %v, want ErrFieldOccupied", err)
		}
	})
}

func TestGame(t *testing.T) {
	t.Run("A game is over when all fields in a diagonal (top-left) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(0)
		_ = g.TakeField(2)
		_ = g.TakeField(4)
		_ = g.TakeField(6)
		_ = g.TakeField(8)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a row (top) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(0)
		_ = g.TakeField(3)
		_ = g.TakeField(1)
		_ = g.TakeField(7)
		_ = g.TakeField(2)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a row (middle) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(3)
		_ = g.TakeField(2)
		_ = g.TakeField(4)
		_ = g.TakeField(7)
		_ = g.TakeField(5)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a row (bottom) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(6)
		_ = g.TakeField(3)
		_ = g.TakeField(7)
		_ = g.TakeField(2)
		_ = g.TakeField(8)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a column (left) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(0)
		_ = g.TakeField(1)
		_ = g.TakeField(3)
		_ = g.TakeField(7)
		_ = g.TakeField(6)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a column (middle) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(1)
		_ = g.TakeField(2)
		_ = g.TakeField(4)
		_ = g.TakeField(5)
		_ = g.TakeField(7)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields in a column (right) are taken by a player", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(2)
		_ = g.TakeField(3)
		_ = g.TakeField(5)
		_ = g.TakeField(6)
		_ = g.TakeField(8)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})

	t.Run("A game is over when all fields are taken", func(t *testing.T) {
		g := NewGame()
		_ = g.TakeField(0)
		_ = g.TakeField(4)
		_ = g.TakeField(2)
		_ = g.TakeField(1)
		_ = g.TakeField(7)
		_ = g.TakeField(5)
		_ = g.TakeField(3)
		_ = g.TakeField(6)
		_ = g.TakeField(8)

		if !g.isGameOver {
			t.Errorf("g.isGameOver should be true")
		}
	})
}

package tictactoe

import "testing"

func TestPlayers(t *testing.T) {
	t.Run("there are two players in the game, X and O", func(t *testing.T) {
		g := NewGame()
		if len(g.players) != 2 {
			t.Errorf("got %d g.players, want 2", len(g.players))
		}
	})
}

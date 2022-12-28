package pgn

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	t.Log(b.Draw())
}

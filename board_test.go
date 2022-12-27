package pgn

import "testing"

func TestNewBoard(t *testing.T) {
	m := map[Square]Piece{}

	b := NewBoard(m)

	t.Log(b.Draw())
}

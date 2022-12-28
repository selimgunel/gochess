package pgn

import "testing"

func TestNewBoard(t *testing.T) {
	b := StartingBoard()

	t.Log(b.Draw())
}

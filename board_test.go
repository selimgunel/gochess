package pgn

import "testing"

func TestNewBoard(t *testing.T) {
	m := map[Square]Piece{}

	swrook := NewSquare(FileA, Rank1)
	m[swrook] = NewPiece(Rook, true)
	b := NewBoard(m)

	t.Log(b.Draw())
}

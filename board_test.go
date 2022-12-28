package pgn

import "testing"

func TestNewBoard(t *testing.T) {
	b := StartingBoard()

	t.Log(b.Draw())
}

func TestBoardFromFen(t *testing.T) {
	b := BoardFromFen("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")

	t.Log(b.Draw())

	t.Log(b.Fen())
}

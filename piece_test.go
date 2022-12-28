package gochess

import "testing"

// TODO: got from notnil/chess
func TestPieceString(t *testing.T) {
	pieces := []struct {
		piece PieceType
		str   string
	}{
		{King, "King"},
		{Queen, "Queen"},
		{Rook, "Rook"},
		{Bishop, "Bishop"},
		{Knight, "Knight"},
		{Pawn, "Pawn"},
	}

	for _, piece := range pieces {
		if piece.piece.String() != piece.str {
			t.Errorf("String version of piece was incorrect.")
		}
	}
}

func TestNewPiece(t *testing.T) {
	testPieces := []struct {
		pieceType PieceType
		color     Color
		pieceName string
		colorName string
	}{
		{King, true, "King", "w"},
		{King, false, "King", "b"},
		{Queen, true, "Queen", "w"},
		{Queen, false, "Queen", "b"},
		{Rook, true, "Rook", "w"},
		{Rook, false, "Rook", "b"},
		{Bishop, true, "Bishop", "w"},
		{Bishop, false, "Bishop", "b"},
		{Knight, true, "Knight", "w"},
		{Knight, false, "Knight", "b"},
		{Pawn, true, "Pawn", "w"},
		{Pawn, false, "Pawn", "b"},
	}

	// test for piece type.
	for _, testPiece := range testPieces {

		piece := NewPiece(testPiece.pieceType, testPiece.color)

		if testPiece.pieceType.String() != piece.PieceType.String() {
			t.Errorf("String version of piece was incorrect.  [want]%s [got]: %s", piece.PieceType, testPiece.pieceType)
		}
	}
}

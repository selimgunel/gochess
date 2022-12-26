package pgn

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
	pieces := []struct {
		piece     PieceType
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

	//test for colors.
	for _, piece := range pieces {
		if piece.color.String() != piece.colorName {
			t.Errorf("String version of color was incorrect.")
		}
	}
	// test for piece type.
	for _, piece := range pieces {
		pName := piece.piece.String()
		if pName != piece.pieceName {
			t.Errorf("String version of piece was incorrect.  [want]%s [got]: %s", pName, piece.pieceName)
		}
	}
}

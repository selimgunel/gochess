package pgn

// PieceType is the type of a piece.
type PieceType int8

type Color bool

const (
	// King represents a king
	King PieceType = iota
	// Queen represents a queen
	Queen
	// Rook represents a rook
	Rook
	// Bishop represents a bishop
	Bishop
	// Knight represents a knight
	Knight
	// Pawn represents a pawn
	Pawn
)

// A piece has a type and color .
type Piece struct {
	PieceType
	Color
}

func NewPiece(typ PieceType, color Color) Piece {

	return Piece{typ, color}
}

func (p PieceType) String() string {
	switch p {
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Rook:
		return "Rook"
	case Bishop:
		return "Bishop"
	case Knight:
		return "Knight"
	case Pawn:
		return "Pawn"
	default:
		return ""
	}
}

func (c Color) String() string {
	if c {
		return "w"
	}
	return "b"
}

func AllPieces() []Piece {

	// whiteKing := Piece{King, true}
	// whiteQueen := Piece{Queen, true}
	// whiteRook := Piece{Rook, true}
	// whiteBishop := Piece{Bishop, true}
	// whiteKnight := Piece{Knight, true}
	// whitePawn := Piece{Knight, true}

	// blackKing := Piece{King, false}
	// blackQueen := Piece{Queen, false}
	// blackRook := Piece{Rook, false}
	// blackBishop := Piece{Bishop, false}
	// blackKnight := Piece{Knight, false}
	// blackPawn := Piece{Knight, false}
	colors := []Color{true, false}
	pieces := make([]Piece, 12)
	pt := []PieceType{King, Queen, Rook, Bishop, Pawn}
	for _, c := range colors {
		for _, t := range pt {
			np := NewPiece(t, c)
			pieces = append(pieces, np)
		}
	}

	return pieces
}

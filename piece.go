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
		return "k"
	case Queen:
		return "q"
	case Rook:
		return "r"
	case Bishop:
		return "b"
	case Knight:
		return "n"
	case Pawn:
		return "p"
	default:
		return ""
	}
}

func (p Piece) Side() string {
	if p.Color {
		return "w"
	}
	return "b"
}

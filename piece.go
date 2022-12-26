package pgn

// PieceType is the type of a piece.
type PieceType int8

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

type Piece struct {
	Type  PieceType
	Color bool
}

func NewPiece(typ PieceType, color bool) Piece {

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

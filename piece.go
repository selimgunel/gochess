package gochess

// PieceType is the type of a piece.
type PieceType int8
type Color bool

const (
	NoPiece PieceType = iota
	// King represents a king
	King
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

func (p Piece) String() string {
	switch p.PieceType {
	case King:
		if p.Color {
			return "K"
		}
		return "k"
	case Queen:
		if p.Color {
			return "Q"
		}
		return "q"
	case Rook:
		if p.Color {
			return "R"
		}
		return "r"
	case Bishop:
		if p.Color {
			return "B"
		}
		return "b"
	case Knight:
		if p.Color {
			return "N"
		}
		return "n"
	case Pawn:
		if p.Color {
			return "P"
		}
		return "p"
	default:
		return ""
	}
}

func (p Piece) Figure() string {
	switch p.PieceType {
	case King:
		if p.Color {
			return "♔"
		}
		return "♚"
	case Queen:
		if p.Color {
			return "♕"
		}
		return "♛"
	case Rook:
		if p.Color {
			return "♖"
		}
		return "♜"
	case Bishop:
		if p.Color {
			return "♗"
		}
		return "♝"
	case Knight:
		if p.Color {
			return "♘"
		}
		return "♞"
	case Pawn:
		if p.Color {
			return "♙"
		}
		return "♟"
	default:
		return ""
	}
}

func pieceFromName(name rune) Piece {
	switch name {
	case 'P':
		return NewPiece(Pawn, true)
	case 'N':
		return NewPiece(Knight, true)
	case 'B':
		return NewPiece(Bishop, true)
	case 'R':
		return NewPiece(Rook, true)
	case 'Q':
		return NewPiece(Queen, true)
	case 'K':
		return NewPiece(King, true)
	case 'p':
		return NewPiece(Pawn, false)
	case 'n':
		return NewPiece(Knight, false)
	case 'b':
		return NewPiece(Bishop, false)
	case 'r':
		return NewPiece(Rook, false)
	case 'q':
		return NewPiece(Queen, false)
	case 'k':
		return NewPiece(King, false)
	}
	return NewPiece(NoPiece, true) //return a white NoPiece
}

func (c Color) String() string {
	if c {
		return "w"
	}
	return "b"
}

func AllPieces() []Piece {

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

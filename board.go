package pgn

import (
	"fmt"
	"strconv"
	"strings"
)

// A Board represents a chess board and its relationship between squares and pieces.
type Board struct {
	whiteKing   Bitboard
	whiteQueen  Bitboard
	whiteRook   Bitboard
	whiteBishop Bitboard
	whiteKnight Bitboard
	whitePawn   Bitboard
	blackKing   Bitboard
	blackQueen  Bitboard
	blackRook   Bitboard
	blackBishop Bitboard
	blackKnight Bitboard
	blackPawn   Bitboard
	whitePieces Bitboard
	blackPieces Bitboard
}

// NewBoard returns a board from a square to piece mapping.
func NewBoard() *Board {
	b := &Board{}

	return b
}

// String implements the fmt.Stringer interface and returns
// a string in the FEN board format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
func (b *Board) String() string {
	fen := ""
	for r := 7; r >= 0; r-- {
		for f := 0; f < 64; f++ {
			sq := NewSquare(File(f), Rank(r))
			p := b.Piece(sq)
			if p.PieceType != NoPiece {
				fen += p.String()
			} else {
				fen += "1"
			}
		}
		if r != 0 {
			fen += "/"
		}
	}
	for i := 8; i > 1; i-- {
		repeatStr := strings.Repeat("1", i)
		countStr := strconv.Itoa(i)
		fen = strings.Replace(fen, repeatStr, countStr, -1)
	}
	return fen
}

// Piece returns the piece for the given square.
func (b *Board) Piece(sq Square) Piece {
	for _, p := range AllPieces() {
		bb := b.bbForPiece(p)
		if bb.Occupied(sq) {
			return p
		}
	}
	return NewPiece(NoPiece, true)
}

func (b *Board) bbForPiece(p Piece) Bitboard {
	switch p.PieceType {
	case King:
		if p.Color {
			return b.whiteKing
		} else {
			return b.blackKing
		}
	case Queen:
		if p.Color {
			return b.whiteQueen
		} else {
			return b.blackQueen
		}
	case Rook:
		if p.Color {
			return b.whiteRook
		} else {
			return b.blackRook
		}
	case Bishop:
		if p.Color {
			return b.whiteBishop
		} else {
			return b.blackBishop
		}
	case Knight:
		if p.Color {
			return b.whiteKnight
		} else {
			return b.blackKnight
		}
	case Pawn:
		if p.Color {
			return b.whitePawn
		} else {
			return b.blackPawn
		}
	}
	return Bitboard(0)
}

// Draw returns visual representation of the board useful for debugging.
func (b *Board) Draw() string {

	s := "\n A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += fmt.Sprint(Rank(r))
		for f := 0; f < len(Files)-1; f++ {
			p := b.PieceAt(SquareOf(File(f), Rank(r)))
			if p.PieceType == NoPiece {
				s += "-"
			} else {
				s += p.Figure()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

var SquareMask = initSquareMask()

func initSquareMask() [64]uint64 {
	var sqm [64]uint64
	for sq := 0; sq < 64; sq++ {
		var b = uint64(1 << sq)
		sqm[sq] = b
	}
	return sqm
}

func (b *Board) PieceAt(sq Square) Piece {
	if sq == NoSquare {
		return NewPiece(NoPiece, true)
	}
	mask := Bitboard(SquareMask[int(sq)])
	if b.blackPieces&mask != 0 {
		if b.blackPawn&mask != 0 {
			return NewPiece(Pawn, false)
		} else if b.blackKnight&mask != 0 {
			return NewPiece(Knight, false)
		} else if b.blackBishop&mask != 0 {
			return NewPiece(Bishop, false)
		} else if b.blackRook&mask != 0 {
			return NewPiece(Rook, false)
		} else if b.blackQueen&mask != 0 {
			return NewPiece(Queen, false)
		} else if b.blackKing&mask != 0 {
			return NewPiece(King, false)
		}
	}

	// It is not black? then it is white
	if b.whitePawn&mask != 0 {
		return NewPiece(Pawn, true)
	} else if b.whiteKnight&mask != 0 {
		return NewPiece(King, true)
	} else if b.whiteBishop&mask != 0 {
		return NewPiece(Bishop, true)
	} else if b.whiteRook&mask != 0 {
		return NewPiece(Rook, true)
	} else if b.whiteQueen&mask != 0 {
		return NewPiece(Queen, true)
	} else if b.whiteKing&mask != 0 {
		return NewPiece(King, true)
	}
	return NewPiece(NoPiece, true)
}

func (b *Board) Clear(square Square, piece Piece) {
	if piece.PieceType == NoPiece {
		return
	}
	mask := Bitboard(SquareMask[int(square)])
	switch piece.PieceType {
	case Pawn:
		switch piece.Color {
		case true:
			b.whitePawn &^= mask
			b.whitePieces &^= mask
		default:
			b.blackPawn &^= mask
			b.blackPieces &^= mask
		}
	case Knight:
		switch piece.Color {
		case true:
			b.whiteKnight &^= mask
			b.whitePieces &^= mask
		default:
			b.blackKnight &^= mask
			b.blackPieces &^= mask
		}
	case Bishop:
		switch piece.Color {
		case true:
			b.whiteBishop &^= mask
			b.whitePieces &^= mask
		default:
			b.blackBishop &^= mask
			b.blackPieces &^= mask
		}
	case Rook:
		switch piece.Color {
		case true:
			b.whiteRook &^= mask
			b.whitePieces &^= mask
		default:
			b.blackRook &^= mask
			b.blackPieces &^= mask
		}
	case Queen:
		switch piece.Color {
		case true:
			b.whiteQueen &^= mask
			b.whitePieces &^= mask
		default:
			b.blackQueen &^= mask
			b.blackPieces &^= mask
		}
	case King:
		switch piece.Color {
		case true:
			b.whiteKing &^= mask
			b.whitePieces &^= mask
		default:
			b.blackKing &^= mask
			b.blackPieces &^= mask
		}
	}
}
func (b *Board) UpdateSquare(sq Square, newPiece Piece, oldPiece Piece) {
	// Remove the piece from source square and add it to destination
	b.Clear(sq, oldPiece)

	mask := Bitboard(SquareMask[int(sq)])
	switch newPiece.PieceType {
	case Pawn:
		switch newPiece.Color {
		case true:
			b.whitePawn |= mask
			b.whitePieces |= mask
		default:
			b.blackPawn |= mask
			b.blackPieces |= mask
		}
	case Knight:
		switch newPiece.Color {
		case true:
			b.whiteKnight |= mask
			b.whitePieces |= mask
		default:
			b.blackKnight |= mask
			b.blackPieces |= mask
		}
	case Bishop:
		switch newPiece.Color {
		case true:
			b.whiteBishop |= mask
			b.whitePieces |= mask
		default:
			b.blackBishop |= mask
			b.blackPieces |= mask
		}
	case Rook:
		switch newPiece.Color {
		case true:
			b.whiteRook |= mask
			b.whitePieces |= mask
		default:
			b.blackRook |= mask
			b.blackPieces |= mask
		}
	case Queen:
		switch newPiece.Color {
		case true:
			b.whiteQueen |= mask
			b.whitePieces |= mask
		default:
			b.blackQueen |= mask
			b.blackPieces |= mask
		}
	case King:
		switch newPiece.Color {
		case true:
			b.whiteKing |= mask
			b.whitePieces |= mask
		default:
			b.blackKing |= mask
			b.blackPieces |= mask
		}
	}
}

func StartingBoard() Board {
	bitboard := Board{}
	noPiece := NewPiece(NoPiece, true)

	bitboard.UpdateSquare(A2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(B2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(C2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(D2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(E2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(F2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(G2, NewPiece(Pawn, true), noPiece)
	bitboard.UpdateSquare(H2, NewPiece(Pawn, true), noPiece)

	bitboard.UpdateSquare(A7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(B7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(C7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(D7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(E7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(F7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(G7, NewPiece(Pawn, false), noPiece)
	bitboard.UpdateSquare(H7, NewPiece(Pawn, false), noPiece)

	bitboard.UpdateSquare(A1, NewPiece(Rook, true), noPiece)
	bitboard.UpdateSquare(B1, NewPiece(Knight, true), noPiece)
	bitboard.UpdateSquare(C1, NewPiece(Bishop, true), noPiece)
	bitboard.UpdateSquare(D1, NewPiece(Queen, true), noPiece)
	bitboard.UpdateSquare(E1, NewPiece(King, true), noPiece)
	bitboard.UpdateSquare(F1, NewPiece(Bishop, true), noPiece)
	bitboard.UpdateSquare(G1, NewPiece(Knight, true), noPiece)
	bitboard.UpdateSquare(H1, NewPiece(Rook, true), noPiece)

	bitboard.UpdateSquare(A8, NewPiece(Rook, false), noPiece)
	bitboard.UpdateSquare(B8, NewPiece(Knight, false), noPiece)
	bitboard.UpdateSquare(C8, NewPiece(Bishop, false), noPiece)
	bitboard.UpdateSquare(D8, NewPiece(Queen, false), noPiece)
	bitboard.UpdateSquare(E8, NewPiece(King, false), noPiece)
	bitboard.UpdateSquare(F8, NewPiece(Bishop, false), noPiece)
	bitboard.UpdateSquare(G8, NewPiece(Knight, false), noPiece)
	bitboard.UpdateSquare(H8, NewPiece(Rook, false), noPiece)

	return bitboard
}

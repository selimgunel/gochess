package pgn

import (
	"fmt"
	"strconv"
	"strings"
)

// A Board represents a chess board and its relationship between squares and pieces.
type Board struct {
	bbWhiteKing   Bitboard
	bbWhiteQueen  Bitboard
	bbWhiteRook   Bitboard
	bbWhiteBishop Bitboard
	bbWhiteKnight Bitboard
	bbWhitePawn   Bitboard
	bbBlackKing   Bitboard
	bbBlackQueen  Bitboard
	bbBlackRook   Bitboard
	bbBlackBishop Bitboard
	bbBlackKnight Bitboard
	bbBlackPawn   Bitboard
	whiteSqs      Bitboard
	blackSqs      Bitboard
	emptySqs      Bitboard
	whiteKingSq   Square
	blackKingSq   Square
}

// NewBoard returns a board from a square to piece mapping.
func NewBoard() *Board {
	b := &Board{}

	return b
}

func (b *Board) setBBForPiece(p Piece, bb Bitboard) {
	switch p.PieceType {
	case King:
		if p.Color {
			b.bbWhiteKing = bb
		} else {
			b.bbBlackKing = bb
		}
	case Queen:
		if p.Color {
			b.bbWhiteQueen = bb
		} else {
			b.bbBlackQueen = bb
		}
	case Rook:
		if p.Color {
			b.bbWhiteRook = bb
		} else {
			b.bbBlackRook = bb
		}
	case Bishop:
		if p.Color {
			b.bbWhiteBishop = bb
		} else {
			b.bbBlackBishop = bb
		}

	case Knight:
		if p.Color {
			b.bbWhiteKnight = bb
		} else {
			b.bbBlackKnight = bb
		}
	case Pawn:
		if p.Color {
			b.bbWhitePawn = bb
		} else {
			b.bbBlackPawn = bb
		}
	case NoPiece:
		fmt.Println("no piece", p)

	default:
		panic("invalid piece")
	}
}

// TODO: understand this better
func (b *Board) calcConvienceBBs() {
	whiteSqs := b.bbWhiteKing | b.bbWhiteQueen | b.bbWhiteRook | b.bbWhiteBishop | b.bbWhiteKnight | b.bbWhitePawn
	blackSqs := b.bbBlackKing | b.bbBlackQueen | b.bbBlackRook | b.bbBlackBishop | b.bbBlackKnight | b.bbBlackPawn
	emptySqs := ^(whiteSqs | blackSqs)
	b.whiteSqs = whiteSqs
	b.blackSqs = blackSqs
	b.emptySqs = emptySqs

	b.whiteKingSq = NoSquare
	b.blackKingSq = NoSquare

	for sq := 0; sq < 64; sq++ {
		sqr := Square(sq)
		if b.bbWhiteKing.Occupied(sqr) {
			b.whiteKingSq = sqr
		} else if b.bbBlackKing.Occupied(sqr) {
			b.blackKingSq = sqr
		}
	}

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
			return b.bbWhiteKing
		} else {
			return b.bbBlackKing
		}
	case Queen:
		if p.Color {
			return b.bbWhiteQueen
		} else {
			return b.bbBlackQueen
		}
	case Rook:
		if p.Color {
			return b.bbWhiteRook
		} else {
			return b.bbBlackRook
		}
	case Bishop:
		if p.Color {
			return b.bbWhiteBishop
		} else {
			return b.bbBlackBishop
		}
	case Knight:
		if p.Color {
			return b.bbWhiteKnight
		} else {
			return b.bbBlackBishop
		}
	case Pawn:
		if p.Color {
			return b.bbWhitePawn
		} else {
			return b.bbBlackPawn
		}
	}
	return Bitboard(0)
}

// Draw returns visual representation of the board useful for debugging.
func (b *Board) Draw() string {
	s := "\n A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += Rank(r).String()
		for f := 0; f < 64; f++ {
			p := b.Piece(NewSquare(File(f), Rank(r)))
			if p.PieceType == NoPiece {
				s += "-"
			} else {
				s += p.String()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

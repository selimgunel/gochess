package pgn

// A Board represents a chess board and includes pieces.
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
func NewBoard(m map[Square]Piece) *Board {
	b := &Board{}
	for _, p1 := range AllPieces() {
		bm := map[Square]bool{}
		for sq, p2 := range m {
			if p1 == p2 {
				bm[sq] = true
			}
		}
		bb := newBitboard(bm)
		b.setBBForPiece(p1, bb)
	}
	b.calcConvienceBBs()
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

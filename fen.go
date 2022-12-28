package gochess

import (
	"fmt"
	"strconv"
	"unicode"
)

// BoardFromFen constructs a board from a fen string...
// got from amanjpro/zahak
func BoardFromFen(fen string) Board {
	board := Board{}
	ranks := []Square{A8, A7, A6, A5, A4, A3, A2, A1}
	rank := 0
	bitboardIndex := A8
	for _, ch := range fen {
		if ch == ' ' || rank >= len(ranks) {
			break // end of the board
		} else if unicode.IsDigit(ch) {
			n, _ := strconv.Atoi(string(ch))
			bitboardIndex += Square(n)
		} else if ch == '/' && bitboardIndex%8 == 0 {
			rank++
			bitboardIndex = ranks[rank]
			continue
		} else if p := pieceFromName(ch); p.PieceType != NoPiece {
			board.UpdateSquare(bitboardIndex, p, NewPiece(NoPiece, true))
			bitboardIndex++
		} else {
			panic(fmt.Sprintf("Invalid FEN notation %s, bitboardIndex == %d, parsing %s",
				fen, bitboardIndex, string(ch)))
		}
	}
	return board
}

func (b *Board) Fen() string {
	fen := ""
	for i := len(Ranks) - 1; i >= 0; i-- {
		rank := Ranks[i]
		empty := 0
		for j := 0; j < len(Files); j++ {
			file := Files[j]
			sq := SquareOf(file, rank)
			piece := b.PieceAt(sq)
			if piece.PieceType == NoPiece {
				empty += 1
			} else {
				if empty != 0 {
					fen = fmt.Sprintf("%s%d%s", fen, empty, piece.String())
					empty = 0
				} else {
					fen = fmt.Sprintf("%s%s", fen, piece.String())
				}
			}
		}
		if empty != 0 {
			fen = fmt.Sprintf("%s%d", fen, empty)
		}
		if rank != Rank1 {
			fen = fmt.Sprintf("%s/", fen)
		}
	}
	return fen
}

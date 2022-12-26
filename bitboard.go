package pgn

import (
	"math/bits"
	"strconv"
	"strings"
)

const sixtyFour = 64

type Bitboard uint64

func (bb Bitboard) String() string {
	s := strconv.FormatUint(uint64(bb), 2)
	var b strings.Builder
	b.Grow(sixtyFour)
	for i := 0; i < sixtyFour-len(s); i++ {
		b.WriteString("0")
	}
	b.WriteString(s)
	return b.String()
}

// Reverse returns a bitboard where the bit order is reversed.
func (bb Bitboard) Reverse() Bitboard {
	return Bitboard(bits.Reverse64(uint64(bb)))
}

func (bb Bitboard) Get() uint64 {
	return uint64(bb)
}

// // String returns a 64 character string of 1s and 0s starting with the most significant bit.
// func (bb Bitboard) StringNe() string {
// 	s := strconv.FormatUint(uint64(bb), 2)
// 	return strings.Repeat("0", numOfSquaresInBoard-len(s)) + s
// }

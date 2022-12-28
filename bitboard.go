package gochess

import (
	"math/bits"
	"strconv"
	"strings"
)

const eight = 8
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

// TODO: This will change, I opt to write a new implementation, this is borrowed from notnil/chess
func (bb Bitboard) Occupied(sq Square) bool {
	return (bits.RotateLeft64(uint64(bb), int(sq)) & 1) == 1
}

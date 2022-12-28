// this file is taken from amanjpro/zahak
package gochess

import (
	"fmt"
)

type Rank int8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

type File int8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

var Ranks = []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}
var Files = []File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

type Square int8

func (sq Square) String() string {
	return sq.Name()
}

const DarkSquares uint64 = 0xAA55AA55AA55AA55

func (sq Square) GetColor() Color {
	return (DarkSquares>>uint64(sq))&1 == 0
}

func (f File) Name() string {
	return string('a' + byte(f))
}

func (r Rank) Name() int {
	return (int(r) + 1)
}

func (s Square) Name() string {
	var fileName = s.File().Name()
	var rankName = s.Rank().Name()
	return fmt.Sprint(fileName, rankName)
}

func (sq Square) File() File {
	return File(sq & 7)
}

func (sq Square) Rank() Rank {
	return Rank(sq / 8)
}

func SquareOf(file File, rank Rank) Square {
	return Square((int(rank) * 8) + int(file))
}

const (
	NoSquare Square = iota - 1
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

var NameToSquareMap = map[string]Square{
	"a1": A1, "a2": A2, "a3": A3, "a4": A4, "a5": A5, "a6": A6, "a7": A7, "a8": A8,
	"b1": B1, "b2": B2, "b3": B3, "b4": B4, "b5": B5, "b6": B6, "b7": B7, "b8": B8,
	"c1": C1, "c2": C2, "c3": C3, "c4": C4, "c5": C5, "c6": C6, "c7": C7, "c8": C8,
	"d1": D1, "d2": D2, "d3": D3, "d4": D4, "d5": D5, "d6": D6, "d7": D7, "d8": D8,
	"e1": E1, "e2": E2, "e3": E3, "e4": E4, "e5": E5, "e6": E6, "e7": E7, "e8": E8,
	"f1": F1, "f2": F2, "f3": F3, "f4": F4, "f5": F5, "f6": F6, "f7": F7, "f8": F8,
	"g1": G1, "g2": G2, "g3": G3, "g4": G4, "g5": G5, "g6": G6, "g7": G7, "g8": G8,
	"h1": H1, "h2": H2, "h3": H3, "h4": H4, "h5": H5, "h6": H6, "h7": H7, "h8": H8,
}

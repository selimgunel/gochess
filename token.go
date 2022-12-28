package gochess

import (
	"fmt"
)

// Token represents a single token in the input stream.
// Name: mnemonic name (numeric).
// Val: string value of the token from the original stream.
// Pos: position - offset from beginning of stream.

type Token struct {
	Name TokenName
	Val  string
	Pos  int
}

type TokenName int

const (
	ERROR TokenName = iota
	EOF

	WS
	TAG
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
	LEFT_ROUND_BRACKET
	RIGHT_ROUND_BRACKET
	QUOTE
	NEWLINE
	COMMENT
	NUMBER

	//Move tokens e4 cx5!
	MOVE
	//Move related tokens. 1.c4 or 8...d5
	TURN_NUMBER
	DOT
	TREE_DOT
	RANK
	FILE

	RESULT
	CHECK //Qd4+

	CAPTURE //Qxd4

)

func (tok Token) String() string {
	return fmt.Sprintf("Token{%s, '%s', %d}", tokenNames[tok.Name], tok.Val, tok.Pos)
}

func makeErrorToken(pos int) Token {
	return Token{ERROR, "", pos}
}

var tokenNames = []string{
	ERROR:    "ERROR",
	EOF:      "EOF",
	COMMENT:  "COMMENT",
	NUMBER:   "NUMBER",
	TAG:      "TAG",
	DOT:      "DOT",
	TREE_DOT: "TREE_DOT",
	MOVE:     "MOVE",
	RESULT:   "RESULT",
}

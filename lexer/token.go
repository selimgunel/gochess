package lexer

// Token represents a lexical token.
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

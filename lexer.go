package pgn

import "unicode/utf8"

// Lexer
//
// Create a new lexer with NewLexer and then call NextToken repeatedly to get
// tokens from the stream. The lexer will return a token with the name EOF when
// done.
type Lexer struct {
	buf string

	// Current rune.
	r rune

	// Position of the current rune in buf.
	rpos int

	// Position of the next rune in buf.
	nextpos int
}

// NewLexer creates a new lexer for the given input.
func NewLexer(buf string) *Lexer {
	lex := Lexer{buf, -1, 0, 0}

	// Prime the lexer by calling .next
	lex.next()
	return &lex
}

// next advances the lexer's internal state to point to the next rune in the
// input.
func (lex *Lexer) next() {
	if lex.nextpos < len(lex.buf) {
		lex.rpos = lex.nextpos
		r, w := rune(lex.buf[lex.nextpos]), 1

		if r >= utf8.RuneSelf {
			r, w = utf8.DecodeRuneInString(lex.buf[lex.nextpos:])
		}

		lex.nextpos += w
		lex.r = r
	} else {
		lex.rpos = len(lex.buf)
		lex.r = -1 // EOF
	}
}

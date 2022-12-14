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

func (lex *Lexer) NextToken() Token {
	// Skip non-tokens like whitespace and check for EOF.
	lex.skipNontokens()
	if lex.r < 0 {
		return Token{EOF, "", lex.nextpos}
	}

	// Is this an operator?
	// if int(lex.r) < len(opTable) {
	// 	if opName := opTable[lex.r]; opName != ERROR {
	// 		if opName == DIVIDE {
	// 			// Special case: '/' may be the start of a comment.
	// 			if lex.peekNextByte() == '/' {
	// 				return lex.scanComment()
	// 			}
	// 		}
	// 		startpos := lex.rpos
	// 		lex.next()
	// 		return Token{opName, string(lex.buf[startpos:lex.rpos]), startpos}
	// 	}
	// }

	// Not an operator. Try other types of tokens.
	// if isAlpha(lex.r) {
	// 	return lex.scanIdentifier()
	// } else if isDigit(lex.r) {
	// 	return lex.scanNumber()
	// } else if lex.r == '"' {
	// 	return lex.scanQuote()
	// }

	if lex.r == '[' {
		return lex.scanTag()
	}

	return makeErrorToken(lex.rpos)
}

func (lex *Lexer) scanTag() Token {
	startpos := lex.rpos
	for lex.r == ']' {
		lex.next()
	}
	return Token{TAG, lex.buf[startpos:lex.rpos], startpos}
}

// peekNextByte returns the next byte in the stream (the one after lex.r).
// Note: a single byte is peeked at - if there's a rune longer than a byte
// there, only its first byte is returned.
func (lex *Lexer) peekNextByte() rune {
	if lex.nextpos < len(lex.buf) {
		return rune(lex.buf[lex.nextpos])
	} else {
		return -1
	}
}

func (lex *Lexer) skipNontokens() {
	for lex.r == ' ' || lex.r == '\t' || lex.r == '\n' || lex.r == '\r' {
		lex.next()
	}
}

func (lex *Lexer) scanIdentifier() Token {
	startpos := lex.rpos
	for isAlpha(lex.r) || isDigit(lex.r) {
		lex.next()
	}
	return Token{IDENTIFIER, lex.buf[startpos:lex.rpos], startpos}
}

func (lex *Lexer) scanNumber() Token {
	startpos := lex.rpos
	for isDigit(lex.r) {
		lex.next()
	}
	return Token{NUMBER, lex.buf[startpos:lex.rpos], startpos}
}

func (lex *Lexer) scanQuote() Token {
	startpos := lex.rpos
	lex.next()
	for lex.r > 0 && lex.r != '"' {
		lex.next()
	}

	if lex.r < 0 {
		return makeErrorToken(startpos)
	} else {
		lex.next()
		return Token{QUOTE, string(lex.buf[startpos:lex.rpos]), startpos}
	}
}

func (lex *Lexer) scanComment() Token {
	startpos := lex.rpos
	lex.next()
	for lex.r > 0 && lex.r != '\n' {
		lex.next()
	}

	tok := Token{COMMENT, string(lex.buf[startpos:lex.rpos]), startpos}
	lex.next()
	return tok
}

func isAlpha(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_' || r == '$'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
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

package pgn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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

func (lex *Lexer) NextToken() Token {
	// Skip non-tokens like whitespace and check for EOF.
	lex.skipNontokens()
	if lex.r < 0 {
		return Token{EOF, "", lex.nextpos}
	}

	if lex.r == '[' {
		lex.next()
		return lex.scanTag()
	} else if isDigit(lex.r) {
		return lex.scanNumber()
	} else if lex.r == '.' {
		return lex.scanDot()
	} else if isMove(lex.r) {
		return lex.scanMove()
	} else if lex.r == '{' {
		return lex.scanComment()
	}

	return makeErrorToken(lex.rpos)
}

func (lex *Lexer) scanDot() Token {

	startpos := lex.rpos

	if lex.peekNextByte() == '.' && lex.peekTwoNextByte() == '.' {
		lex.next()
		lex.next()
		lex.next()
		return Token{TREE_DOT, lex.buf[startpos:lex.rpos], startpos}
	}
	lex.next()
	return Token{DOT, lex.buf[startpos:lex.rpos], startpos}

}

func (lex *Lexer) scanMove() Token {
	startpos := lex.rpos

	for isMove(lex.r) {
		lex.next()
	}

	return Token{MOVE, lex.buf[startpos:lex.rpos], startpos}
}

func (lex *Lexer) scanTag() Token {
	startpos := lex.rpos

	for lex.r != ']' {
		lex.next()
	}
	defer lex.next()
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

func (lex *Lexer) peekTwoNextByte() rune {
	if lex.nextpos < len(lex.buf) {
		return rune(lex.buf[lex.nextpos+1])
	} else {
		return -1
	}
}

func (lex *Lexer) skipNontokens() {
	for lex.r == ' ' || lex.r == '\t' || lex.r == '\n' || lex.r == '\r' {
		lex.next()
	}
}

func (lex *Lexer) scanNumber() Token {
	startpos := lex.rpos

	if lex.r == '0' || lex.r == '1' {

		b := lex.peekNextByte()
		fmt.Println("result part")
		if b == '-' || b == '/' {
			for isResult(lex.r) {
				lex.next()
			}
			return Token{RESULT, lex.buf[startpos:lex.rpos], startpos}
		}
	}

	for isDigit(lex.r) {
		lex.next()
	}
	return Token{NUMBER, lex.buf[startpos:lex.rpos], startpos}
}

func isResult(r rune) bool {

	return strings.ContainsRune("12/0-", r)
}
func (lex *Lexer) scanComment() Token {
	startpos := lex.rpos
	lex.next()
	for lex.r != '}' {
		lex.next()
	}

	tok := Token{COMMENT, string(lex.buf[startpos:lex.rpos]), startpos}
	lex.next()
	return tok
}

func isMove(r rune) bool {
	return strings.ContainsRune("KNRQBabcdefgh0123456789O-+x!?", r)
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

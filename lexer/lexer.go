package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"
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
	LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing comment"
	LEXER_ERROR_MOVE_NUMBER_EXPECTED  string = "A digit expected"
)

const (
	ERROR TokenName = iota
	EOF

	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
	LEFT_ROUND_BRACKET
	RIGHT_ROUND_BRACKET
	NEWLINE
	COMMENT

	//Move tokens e4 cx5!
	WHITE
	BLACK
	//Move related tokens. 1.c4 or 8...d5
	TURN_NUMBER
	DOT
	TREE_DOT
	RANK
	FILE

	CHECK //Qd4+

	CAPTURE //Qxd4

)
const eof = -1

var tokenNames = []string{
	ERROR:               "ERROR",
	EOF:                 "EOF",
	LEFT_CURLY_BRACKET:  "LEFT_CURLY_BRACKET",
	RIGHT_CURLY_BRACKET: "RIGHT_CURLY_BRACKET",
	LEFT_ROUND_BRACKET:  "LEFT_ROUND_BRACKET",
	RIGHT_ROUND_BRACKET: "RIGHT_ROUND_BRACKET",
	NEWLINE:             "NEWLINE",
	COMMENT:             "COMMENT",
	WHITE:               "WHITE",
	BLACK:               "BLACK",
	TURN_NUMBER:         "TURN_NUMBER",
	DOT:                 "DOT",
	TREE_DOT:            "TREE_DOT",
	RANK:                "RANK",
	FILE:                "FILE",
	CHECK:               "CHECK",
	CAPTURE:             "CAPTURE",
}

func (tok Token) String() string {
	return fmt.Sprintf("Token{%s, '%s', %d}", tokenNames[tok.Name], tok.Val, tok.Pos)
}

type Lexer struct {
	input  string     // the string being scanned
	start  int        // start position of this tokens
	pos    int        // current position of the input
	width  int        // width of last rune read from input
	tokens chan Token // channel of scanned tokens
}

type stateFn func(*Lexer) stateFn

// Lex creates a new Lexer
func Lex(input string) *Lexer {
	l := &Lexer{
		input:  input,
		tokens: make(chan Token),
	}
	go l.run()
	return l
}

// NextToken returns the next item from the input. The Lexer has to be
// drained (all tokens received until itemEOF or itemError) - otherwise
// the Lexer goroutine will leak.
func (l *Lexer) NextToken() Token {
	return <-l.tokens
}

// run runs the lexer - should be run in a separate goroutine.
func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens) // no more tokens will be delivered
}

func (l *Lexer) emit(name TokenName) {
	l.tokens <- Token{
		Name: name,
		Val:  l.input[l.start:l.pos],
		Pos:  l.start,
	}
	l.start = l.pos
}

// next advances to the next rune in input and returns it
func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. Can be called only once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next run in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// errorf returns an error token and terminates the scan by passing back
// a nil pointer that will be the next state.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- Token{
		Name: ERROR,
		Val:  fmt.Sprintf(format, args...),
		Pos:  l.pos,
	}
	return nil
}

// isAlpha reports whether r is an alphabetic or underscore.
func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// isDigit reports whether r is a digit.
func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func lexText(l *Lexer) stateFn {

	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(EOF)
			return nil
		case r == ' ' || r == '\t' || r == '\n' || r == '\r':
			l.ignore()
		case int(r) < len(opTable) && opTable[r] != ERROR:
			op := opTable[r]
			if op == DIVIDE && l.peek() == '/' {
				return lexComment
			}
			l.emit(op)
		case isAlpha(r):
			l.backup()
			return lexIdentifier
		case isDigit(r):
			l.backup()
			return lexNumber
		case r == '"':
			return lexQuote
		}
	}
}

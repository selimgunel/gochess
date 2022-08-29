package lexer

import (
	"fmt"
	"strings"
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

	TAG
	LEFT_CURLY_BRACKET
	RIGHT_CURLY_BRACKET
	LEFT_ROUND_BRACKET
	RIGHT_ROUND_BRACKET
	NEWLINE
	COMMENT
	NUMBER
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
	TAG:                 "TAG",
	LEFT_CURLY_BRACKET:  "LEFT_CURLY_BRACKET",
	RIGHT_CURLY_BRACKET: "RIGHT_CURLY_BRACKET",
	LEFT_ROUND_BRACKET:  "LEFT_ROUND_BRACKET",
	RIGHT_ROUND_BRACKET: "RIGHT_ROUND_BRACKET",
	NEWLINE:             "NEWLINE",
	COMMENT:             "COMMENT",
	WHITE:               "WHITE",
	BLACK:               "BLACK",
	NUMBER:              "NUMBER",
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
	fmt.Println("should read")
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

// isDigit reports whether r is a digit.
func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// Operator table for lookups.
var opTable = [...]TokenName{
	'{': LEFT_CURLY_BRACKET,
	'}': RIGHT_CURLY_BRACKET,
	'(': LEFT_ROUND_BRACKET,
	')': RIGHT_ROUND_BRACKET,
	'.': DOT,
}

func lexComment(l *Lexer) stateFn {
	r := l.next()
	for r != eof && r != '\n' && r != '}' {
		r = l.next()
	}
	l.ignore()
	return lexText
}

func lexNumber(l *Lexer) stateFn {
	for {
		r := l.next()

		if r == '.' {
			l.backup()
			break
		}
	}
	l.emit(TURN_NUMBER)
	return lexText
}

func lexMove(l *Lexer) stateFn {
	for {

		r := l.next()
		if !isMove(r) {
			break
		}
	}
	l.backup()
	l.emit(BLACK)
	return lexText
}

func lexTag(l *Lexer) stateFn {

out:
	for {
		r := l.next()
		if r != ']' {
			l.next()
		} else {
			break out
		}

	}

	l.emit(TAG)
	return lexText
}

// isMove reports whether r has a valid letter.
func isMoveStart(r rune) bool {
	return strings.ContainsRune("KQRBNabcdefghO", r)
}

// isMove reports whether r has a valid letter.
func isMove(r rune) bool {
	return strings.ContainsRune("KQRBNabcdefgh0123456789x!?/-O+", r)
}

func lexText(l *Lexer) stateFn {

	for {
		switch r := l.next(); {
		case r == eof:
			l.emit(EOF)
		case r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '.':
			l.ignore()
		case r == '{':
			l.ignore()
			return lexComment

		case isDigit(r):
			l.backup()
			return lexNumber
		case isMoveStart(r):
			l.backup()
			return lexMove
		case r == '[':
			l.backup()
			return lexTag
		default:
			fmt.Printf("shouldn't reach here\n")
			l.emit(ERROR)
			return nil

		}
	}
}

func TokenizeAllAppend(input string) []Token {
	var tokens []Token
	l := Lex(input)
	for {
		token := l.NextToken()
		fmt.Printf("tok: %s\n", token)
		tokens = append(tokens, token)
		if token.Name == EOF || token.Name == ERROR {
			break
		}
	}
	return tokens
}

func TokenizeAllPrealloc(input string) []Token {
	tokens := make([]Token, 0, 200000)
	l := Lex(input)
	for {
		token := l.NextToken()
		tokens = append(tokens, token)
		if token.Name == EOF || token.Name == ERROR {
			break
		}
	}
	return tokens
}

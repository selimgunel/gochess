package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Token represents a single token in the input stream.
// Name: mnemonic name (numeric).
// Val: string value of the token from the original stream.
// Pos: position - offset from beginning of stream.

const (
	LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing comment"
	LEXER_ERROR_MOVE_NUMBER_EXPECTED  string = "A digit expected"
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
	MOVE:                "MOVE",

	NUMBER:      "NUMBER",
	TURN_NUMBER: "TURN_NUMBER",
	DOT:         "DOT",
	TREE_DOT:    "TREE_DOT",
	RANK:        "RANK",
	FILE:        "FILE",
	CHECK:       "CHECK",
	CAPTURE:     "CAPTURE",
	WS:          "WS",
	RESULT:      "RESULT",
}

func (tok Token) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "Token{%s, '%s', %d}", tokenNames[tok.Name], tok.Val, tok.Pos)
	return s.String()
}

type Lexer struct {
	r   *bufio.Reader
	pos int
}

// NewLexer returns a new instance of Scanner.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: bufio.NewReader(r),
		pos: 0,
	}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Lexer) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	s.pos++
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Lexer) unread() { _ = s.r.UnreadRune(); s.pos-- }

// Scan returns the next token and literal value.
func (s *Lexer) Scan() (tok Token) {
	// Read the next rune.
	ch := s.read()

	//check if it is a digit.
	if isDigit(ch) {
		s.unread()
		return s.readTurnNumber()
	}

	switch ch {
	case eof:
		return Token{Name: EOF, Val: "", Pos: s.pos}
	case '[':
		s.unread()
		return s.readTag()
	case '\n':
		return Token{Name: NEWLINE, Val: `\n`, Pos: s.pos}
	case '{':
		s.unread()
		return s.readComment()
	default:
		s.unread()
		return s.readMove()

	}

}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Lexer) readTurnNumber() (tok Token) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if ch == '.' {

			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return Token{Name: TURN_NUMBER, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

// readTag consumes the current rune and all contiguous whitespace.
func (s *Lexer) readMove() (tok Token) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if isMove(ch) {
			buf.WriteRune(ch)
		} else {
			break
		}

	}

	return Token{Name: MOVE, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}

}

func (s *Lexer) readTag() (tok Token) {
	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if ch == ']' {
			break

		} else if ch == '[' {
			s.unread()
			s.r.Discard(1)
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Name: TAG, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

func (s *Lexer) readComment() (tok Token) {

	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if ch == '}' {

			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Name: COMMENT, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// isDigit returns true if the rune is a digit.
func isResultChar(ch rune) bool { return strings.ContainsRune("12/-", ch) }

func isMove(ch rune) bool {
	return strings.ContainsRune("KNRQBabcdefgh0123456789O-+x!?", ch)
}

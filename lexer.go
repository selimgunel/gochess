package parser

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const (
	LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing comment"
	LEXER_ERROR_MOVE_NUMBER_EXPECTED  string = "A digit expected"
)

const eof = -1

type Lexer struct {
	r   *bufio.Reader
	pos int
}

// NewLexer returns a new instance of Lexer.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: bufio.NewReader(r),
		pos: 0,
	}
}

// read reads the next rune from the bufferred reader.
// Returns the -1 if an error occurs.
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
	case '(':
		s.unread()
		return s.readRoundComment()
	case '$':
		s.unread()
		return s.readPin()

	default:
		s.unread()
		return s.readMove()

	}

}

// readPin reads placeholders e.g $12 .
func (s *Lexer) readPin() (tok Token) {

	s.r.Discard(1)
	for {
		if ch := s.read(); ch == eof {
			panic("it shouldn't reach here")
		} else if isDigit(ch) {
			s.r.Discard(1)
		} else {
			break
		}
	}
	return Token{}
}

// readTurnNumber reads turn number.
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
		} else if ch == '/' || ch == '-' {
			//TODO: that is for parsing the result clause, it doesn't work now.
			p, err := s.r.Peek(1)
			if err != nil {
				return Token{Name: ERROR, Val: "peek error", Pos: s.pos - len(buf.Bytes())}
			}
			if !isResultChar(rune(p[0])) {
				break
			}
			buf.WriteRune(ch)

		} else {
			buf.WriteRune(ch)
		}
	}
	return Token{Name: TURN_NUMBER, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

// readMove consumes the current move.
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

// readMove consumes the current tag.
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

	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if ch == '}' {

			break
		} else if ch == '{' {
			s.unread()
			s.r.Discard(1)
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Name: COMMENT, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

func (s *Lexer) readRoundComment() (tok Token) {

	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
		} else if ch == ')' {

			break
		} else if ch == '(' {
			s.unread()
			s.r.Discard(1)
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{Name: COMMENT, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// isResultChar returns true if the rune is a valid result char. TODO: it is not useful now.
func isResultChar(ch rune) bool { return strings.ContainsRune("12/-", ch) }

// isMove returns true if the rune is a valid move character.
func isMove(ch rune) bool {
	return strings.ContainsRune("KNRQBabcdefgh0123456789O-+x!?", ch)
}

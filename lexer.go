package pgn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// item represents a token or text string returned from the scanner.
type item struct {
	typ  itemType // The type of this item.
	pos  Pos      // The starting position, in bytes, of this item in the input string.
	val  string   // The value of this item.
	line int      // The line number at the start of this item.
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError itemType = iota // error occurred; value is text of error
	itemRightBracket
	itemLeftBracket
	itemSpace
	itemComment // comment text
	itemNumber  //simple number
	itemEOF
)

//Pos represents a byte position in the input.

type Pos int

// lexer holds the state of the scanner.

type lexer struct {
	input string // the string being scanned

	pos   Pos  // current position in the input
	start Pos  // start position of this item
	atEOF bool // we have hit the end of input and returned eof

	line      int       // 1+number of newlines seen
	startLine int       // start line of this item
	items     chan item // item to return to parser

}

const eof = -1

// Trimming spaces.
// If the action begins "{{- " rather than "{{", then all space/tab/newlines
// preceding the action are trimmed; conversely if it ends " -}}" the
// leading spaces are trimmed. This is done entirely in the lexer; the
// parser never sees it happen. We require an ASCII space (' ', \t, \r, \n)
// to be present to avoid ambiguity with things like "{{-3}}". It reads
// better with the space present anyway. For simplicity, only ASCII
// does the job.
const (
	spaceChars    = " \t\r\n"  // These are the space characters defined by Go itself.
	trimMarker    = '-'        // Attached to left/right delimiter, trims trailing spaces from preceding/following text.
	trimMarkerLen = Pos(1 + 1) // marker plus space before or after
)

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.atEOF = true
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += Pos(w)
	if r == '\n' {
		l.line++
	}
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune.
func (l *lexer) backup() {
	if !l.atEOF && l.pos > 0 {
		r, w := utf8.DecodeLastRuneInString(l.input[:l.pos])
		l.pos -= Pos(w)
		// Correct newline count.
		if r == '\n' {
			l.line--
		}
	}
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.startLine}
	l.start = l.pos
	l.startLine = l.line
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.line += strings.Count(l.input[l.start:l.pos], "\n")
	l.start = l.pos
	l.startLine = l.line
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...any) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.startLine}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	return <-l.items
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}

// lex creates a new scanner for the input string.
func lex(input, left, right string, emitComment, breakOK, continueOK bool) *lexer {

	l := &lexer{

		input: input,

		items:     make(chan item),
		line:      1,
		startLine: 1,
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// func (s *Lexer) readRoundComment() (tok Token) {

// 	var buf bytes.Buffer

// 	for {
// 		if ch := s.read(); ch == eof {
// 			return Token{Name: ERROR, Val: "eof reached", Pos: s.pos}
// 		} else if ch == ')' {

// 			break
// 		} else if ch == '(' {
// 			s.unread()
// 			s.r.Discard(1)
// 		} else {
// 			buf.WriteRune(ch)
// 		}
// 	}

// 	return Token{Name: COMMENT, Val: buf.String(), Pos: s.pos - len(buf.Bytes())}
// }

// // isDigit returns true if the rune is a digit.
// func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// // isResultChar returns true if the rune is a valid result char. TODO: it is not useful now.
// func isResultChar(ch rune) bool { return strings.ContainsRune("12/-", ch) }

// // isMove returns true if the rune is a valid move character.
// func isMove(ch rune) bool {
// 	return strings.ContainsRune("KNRQBabcdefgh0123456789O-+x!?", ch)
// }

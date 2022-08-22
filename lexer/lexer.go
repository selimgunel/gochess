package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const EOF rune = 0

const LEFT_CURLY_BRACKET string = "{"
const RIGHT_CURLY_BRACKET string = "}"
const LEFT_ROUND_BRACKET string = "("
const RIGHT_ROUND_BRACKET string = ")"

const EQUAL_SIGN string = "="
const NEWLINE string = "\n"

const (
	LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing section bracket"
)

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_LEFT_CURLY_BRACKET
	TOKEN_RIGHT_CURLY_BRACKET
	TOKEN_LEFT_ROUND_BRACKET
	TOKEN_RIGHT_ROUND_BRACKET
	TOKEN_NEWLINE
	TOKEN_COMMENT

	//Move related tokens. 1.c4 or 8...d5
	TOKEN_TURN_NUMBER
	TOKEN_DOT
	TOKEN_TREE_DOT
	TOKEN_RANK
	TOKEN_FILE

	TOKEN_CHECK //Qd4+

	TOKEN_CAPTURE //Qxd4

)

type LexFn func(*Lexer) LexFn

type Lexer struct {
	Name   string
	Input  string
	Tokens chan Token
	State  LexFn

	Start int
	Pos   int
	Width int
}

/*
l lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), LEFT_CURLY_BRACKET) {
		return LexLeftCurlyBracket
	}
	return LexReadMove

}

/*
l lexer function emits a TOKEN_COMMENT then returns
the lexer for value.
*/
func LexLeftCurlyBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(LEFT_CURLY_BRACKET)
	lexer.Emit(TOKEN_LEFT_CURLY_BRACKET)
	return LexComment
}

/*
l lexer function emits a TOKEN_COMMENT then returns
the lexer for value.
*/
func LexComment(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(LEXER_ERROR_MISSING_RIGHT_BRACKET)
		}

		if strings.HasPrefix(lexer.InputToEnd(), RIGHT_CURLY_BRACKET) {
			lexer.Emit(TOKEN_COMMENT)
			return LexRightCurlyBracket
		}

		lexer.Inc()
	}
}

/*
l lexer function emits a TOKEN_COMMENT then returns
the lexer for value.
*/
func LexRightCurlyBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(RIGHT_CURLY_BRACKET)
	lexer.Emit(TOKEN_RIGHT_CURLY_BRACKET)
	return LexReadMove
}

/*
l lexer function emits a TOKEN_COMMENT then returns
the lexer for value.
*/
func LexReadMove(l *Lexer) LexFn {

	for {

		if unicode.IsDigit(l.Peek()) {
			l.Emit(TOKEN_TURN_NUMBER)
			return
		}

		l.Inc()
	}
}

/*
Start a new lexer with a given input string. l returns the
instance of the lexer and a channel of tokens. Reading l stream
is the way to parse a given input and perform processing.
*/
func BeginLexing(name, input string) *Lexer {
	l := &Lexer{
		Name:   name,
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan Token, 3),
	}

	return l
}

/*
Backup to the beginning of the last read token.
*/
func (l *Lexer) Backup() {
	l.Pos -= l.Width
}

/*
Returns a slice of the current input from the current lexer start position
to the current position.
*/
func (l *Lexer) CurrentInput() string {
	return l.Input[l.Start:l.Pos]
}

/*
Decrement the position
*/
func (l *Lexer) Dec() {
	l.Pos--
}

/*
Puts a token onto the token channel. The value of l token is
read from the input based on the current lexer position.
*/
func (l *Lexer) Emit(tokenType TokenType) {
	l.Tokens <- Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

/*
Returns a token with error information.
*/
func (l *Lexer) Errorf(format string, args ...interface{}) LexFn {
	l.Tokens <- Token{
		Type:  TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

/*
Ignores the current token by setting the lexer's start
position to the current reading position.
*/
func (l *Lexer) Ignore() {
	l.Start = l.Pos
}

/*
Increment the position
*/
func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Emit(TOKEN_EOF)
	}
}

/*
Return a slice of the input from the current lexer position
to the end of the input string.
*/
func (l *Lexer) InputToEnd() string {
	return l.Input[l.Pos:]
}

/*
Returns the true/false if the lexer is at the end of the
input stream.
*/
func (l *Lexer) IsEOF() bool {
	return l.Pos >= len(l.Input)
}

/*
Returns true/false if then next character is whitespace
*/
func (l *Lexer) IsWhitespace() bool {
	ch, _ := utf8.DecodeRuneInString(l.Input[l.Pos:])
	return unicode.IsSpace(ch)
}

/*
Reads the next rune (character) from the input stream
and advances the lexer position.
*/
func (l *Lexer) Next() rune {
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Width = 0
		return EOF
	}

	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])

	l.Width = width
	l.Pos += l.Width
	return result
}

/*
Return the next token from the channel
*/
func (l *Lexer) NextToken() Token {
	for {
		select {
		case token := <-l.Tokens:
			return token
		default:
			l.State = l.State(l)
		}
	}

}

/*
Returns the next rune in the stream, then puts the lexer
position back. Basically reads the next rune without consuming
it.
*/
func (l *Lexer) Peek() rune {
	rune := l.Next()
	l.Backup()
	return rune
}

/*
Starts the lexical analysis and feeding tokens into the
token channel.
*/
func (l *Lexer) Run() {
	for state := LexBegin; state != nil; {
		state = state(l)
	}

	l.Shutdown()
}

/*
Shuts down the token stream
*/
func (l *Lexer) Shutdown() {
	close(l.Tokens)
}

/*
Skips whitespace until we get something meaningful.
*/
func (l *Lexer) SkipWhitespace() {
	for {
		ch := l.Next()

		if !unicode.IsSpace(ch) {
			l.Dec()
			break
		}

		if ch == EOF {
			l.Emit(TOKEN_EOF)
			break
		}
	}
}

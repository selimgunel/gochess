package lexer_test

import (
	"strings"
	"testing"

	"github.com/narslan/schach/pgnparser/lexer"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s        string
		expected lexer.TokenName
		tok      lexer.Token
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, expected: lexer.EOF, tok: lexer.Token{Name: lexer.EOF, Val: "", Pos: 0}},
		{s: `[Event "?"]`, expected: lexer.TAG, tok: lexer.Token{Name: lexer.TAG, Val: `[Event "?"]`, Pos: 0}},
	}

	for i, tt := range tests {
		s := lexer.NewLexer(strings.NewReader(tt.s))
		tok := s.Scan()
		if tt.tok.Name != tt.expected {
			t.Errorf("%d. token mismatch: exp=%q got=%q", i, tt.expected, tok.Name)
		}
		if tt.s != tok.Val {
			t.Errorf("%d. token fail: exp=%q got=%q", i, tt.s, tok.Val)
		}
	}
}

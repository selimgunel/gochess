package parser_test

import (
	"strings"
	"testing"

	"github.com/narslan/schach/pgnparser/parser"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s        string
		expected parser.TokenName
		tok      parser.Token
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, expected: parser.EOF, tok: parser.Token{Name: parser.EOF, Val: "", Pos: 0}},
		{s: `[Event "?"]`, expected: parser.TAG, tok: parser.Token{Name: parser.TAG, Val: `[Event "?"]`, Pos: 0}},
	}

	for i, tt := range tests {
		s := parser.NewLexer(strings.NewReader(tt.s))
		tok := s.Scan()
		if tt.tok.Name != tt.expected {
			t.Errorf("%d. token mismatch: exp=%q got=%q", i, tt.expected, tok.Name)
		}
		if tt.s != tok.Val {
			t.Errorf("%d. token fail: exp=%q got=%q", i, tt.s, tok.Val)
		}
	}
}

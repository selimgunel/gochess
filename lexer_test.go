package pgn_test

import (
	"strings"
	"testing"

	lexer "github.com/narslan/pgn"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s            string
		expectedName lexer.TokenName
		expectedVal  string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, expectedName: lexer.EOF, expectedVal: ""},
		{s: `[Event "?"]`, expectedName: lexer.TAG, expectedVal: `Event "?"`},
		{s: `[White "?"]`, expectedName: lexer.TAG, expectedVal: `White "?"`},
		{s: `[Black "?"]`, expectedName: lexer.TAG, expectedVal: `Black "?"`},
		{s: `{}`, expectedName: lexer.COMMENT, expectedVal: ""},
		{s: `1.`, expectedName: lexer.TURN_NUMBER, expectedVal: "1"},
	}

	for i, tt := range tests {
		s := lexer.NewLexer(strings.NewReader(tt.s))
		tok := s.Scan()
		if tok.Name != tt.expectedName {
			t.Errorf("%d. token mismatch: exp=%q got=%q", i, tt.expectedName, tok.Name)
		}
		if tok.Val != tt.expectedVal {
			t.Errorf("%d. token fail: exp=%q got=%q", i, tt.s, tok.Val)
		}
	}
}

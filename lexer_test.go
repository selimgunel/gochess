package parser_test

import (
	"strings"
	"testing"

	parser "github.com/narslan/pgn"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s            string
		expectedName parser.TokenName
		expectedVal  string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, expectedName: parser.EOF, expectedVal: ""},
		{s: `[Event "?"]`, expectedName: parser.TAG, expectedVal: `Event "?"`},
		{s: `{}`, expectedName: parser.COMMENT, expectedVal: ""},
		{s: `1.`, expectedName: parser.TURN_NUMBER, expectedVal: "1"},
	}

	for i, tt := range tests {
		s := parser.NewLexer(strings.NewReader(tt.s))
		tok := s.Scan()
		if tok.Name != tt.expectedName {
			t.Errorf("%d. token mismatch: exp=%q got=%q", i, tt.expectedName, tok.Name)
		}
		if tok.Val != tt.expectedVal {
			t.Errorf("%d. token fail: exp=%q got=%q", i, tt.s, tok.Val)
		}
	}
}

func TestScanner_MultipleTokens(t *testing.T) {
	var tests = []struct {
		s            string
		expectedName parser.TokenName
		expectedVal  string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, expectedName: parser.EOF, expectedVal: ""},
		{s: `[Event "?"]`, expectedName: parser.TAG, expectedVal: `Event "?"`},
		{s: `{}`, expectedName: parser.COMMENT, expectedVal: ""},
		{s: `1.`, expectedName: parser.TURN_NUMBER, expectedVal: "1"},
	}

	for i, tt := range tests {
		s := parser.NewLexer(strings.NewReader(tt.s))
		tok := s.Scan()
		if tok.Name != tt.expectedName {
			t.Errorf("%d. token mismatch: exp=%q got=%q", i, tt.expectedName, tok.Name)
		}
		if tok.Val != tt.expectedVal {
			t.Errorf("%d. token fail: exp=%q got=%q", i, tt.s, tok.Val)
		}
	}
}

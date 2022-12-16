package pgn

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type lexTest struct {
	name   string
	input  string
	tokens []Token
}

func mkItem(name TokenName, val string) Token {
	return Token{
		Name: name,
		Val:  val,
	}
}

var (
	tEOF = mkItem(EOF, "")

	tTag      = mkItem(TAG, `Event "test"`)
	tTagOther = mkItem(TAG, `AnotherEvent "test"`)
	tNumber   = mkItem(NUMBER, `1`)
	tDot      = mkItem(DOT, `.`)
	ttreeDot  = mkItem(TREE_DOT, `...`)
	tMove     = mkItem(MOVE, `Nb5`)
)

var twoLineTags string = `[Event "test"]
[AnotherEvent "test"]
`

var twoLineTagsNumber string = `[Event "test"]
[AnotherEvent "test"] 1
`

var lexTests = []lexTest{
	{"empty", "", []Token{tEOF}},
	{"tag", `[Event "test"]`, []Token{tTag, tEOF}},
	{"multiline", twoLineTags, []Token{tTag, tTagOther, tEOF}},
	{"multiline number", twoLineTagsNumber, []Token{tTag, tTagOther, tNumber, tEOF}},
	{"multiline number with dot", twoLineTagsNumber + ".", []Token{tTag, tTagOther, tNumber, tDot, tEOF}},
	{"multiline number with tree dot", twoLineTagsNumber + " ...", []Token{tTag, tTagOther, tNumber, ttreeDot, tEOF}},

	{"read move", twoLineTags + " 1. Nb5", []Token{tTag, tTagOther, tNumber, tDot, tMove, tEOF}},
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest, tb testing.TB) (items []Token) {

	lex := NewLexer(t.input)
	var toks []Token

	for {
		tok := lex.NextToken()
		toks = append(toks, tok)
		if tok.Name == EOF {
			break
		}

	}
	return toks
}

func TestLex(t *testing.T) {
	for _, test := range lexTests {
		t.Log(test.name, "collecting")
		items := collect(&test, t)
		if !equal(items, test.tokens) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, items, test.tokens)

		}
		t.Log(test.name, "OK")
	}
}

func equal(i1, i2 []Token) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].Val != i2[k].Val {
			return false
		}

	}
	return true
}

func TestParse(t *testing.T) {

	//f, err := os.ReadFile("testdata/s1.pgn")
	f, err := os.ReadFile("pgn/counter-vs-zahak.pgn")
	checkErr(err, t)
	moves, _ := Parse(string(f))

	var b strings.Builder

	for i, v := range moves {
		if i%2 == 0 {
			fmt.Fprintf(&b, "%d. ", i/2+1)
		}
		fmt.Fprintf(&b, "%s ", v)
	}
	t.Log(b.String())
}

func checkErr(err error, tb testing.TB) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

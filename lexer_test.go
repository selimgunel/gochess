package pgn

import "testing"

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
)

var twoLineTags string = `[Event "test"]
[AnotherEvent "test"]
`
var lexTests = []lexTest{
	{"empty", "", []Token{tEOF}},
	{"tag", `[Event "test"]`, []Token{tTag, tEOF}},
	{"multiline", twoLineTags, []Token{tTag, tTagOther, tEOF}},
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

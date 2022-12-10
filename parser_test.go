package pgn

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {

	f, err := os.Open("testdata/two.pgn")
	checkErr(err, t)
	tags, moves, err := Split(f)
	checkErr(err, t)
	for _, v := range tags {
		t.Logf("%v\n", v)
	}
	for _, v := range moves {
		t.Logf("%v\n", v)
	}

	Parse(tags, moves)

}

func checkErr(err error, tb testing.TB) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

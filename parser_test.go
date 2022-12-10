package pgn

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {

	f, err := os.Open("testdata/small.pgn")
	checkErr(err, t)
	games, err := Parse(f)
	checkErr(err, t)
	for _, v := range games {
		t.Logf("%v\n", v.Moves)
	}
	for _, v := range games {
		t.Logf("%v\n", v.Tags)
	}
}

func checkErr(err error, tb testing.TB) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

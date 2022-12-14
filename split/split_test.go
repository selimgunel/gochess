package pgn

import (
	"os"
	"testing"
)

func TestSplit(t *testing.T) {

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

	//Parse(tags, moves)

}

func TestSplitPoint(t *testing.T) {

	f, err := os.Open("testdata/two.pgn")
	checkErr(err, t)
	sps, err := SplitPoints(f)
	checkErr(err, t)
	for _, v := range sps {
		t.Logf("%v\n", v)
	}

}

// func TestSplitAndSave(t *testing.T) {

// 	err := SplitAndSave("testdata/two.pgn")
// 	checkErr(err, t)

// }

// TODO: that is for configuration with bits
func TestProcstr(t *testing.T) {

	procstr("HELLO PEOPLE!", LOWER|REV|CAP)

}

func checkErr(err error, tb testing.TB) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

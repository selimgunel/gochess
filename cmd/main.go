package main

import (
	"os"

	"github.com/narslan/schach/pgnparser/parser"
)

func main() {

	f, err := os.Open("/home/nevroz/go/src/github.com/narslan/schach/pgnparser/pgn/bali02.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	parser.Parse(f)

}

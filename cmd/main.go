package main

import (
	"fmt"
	"os"

	parser "github.com/narslan/pgn"
)

func main() {

	f, err := os.Open("/home/nevroz/go/src/github.com/narslan/pgn/pgn/polgar.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	games := parser.Parse(f)
	for _, v := range games {
		fmt.Printf("%v\n", v.Moves)
	}

}

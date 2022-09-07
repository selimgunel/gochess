package main

import (
	parser "github.com/narslan/pgn"
)

func main() {

	//pgn parse.
	// f, err := os.Open("/home/nevroz/go/src/github.com/narslan/pgn/pgn/famous_games.pgn")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// games := parser.Parse(f)
	// for _, v := range games {
	// 	fmt.Printf("%v\n", v.Moves)
	// }

	//move ...
	parser.PosFromStart("/home/nevroz/go/src/github.com/narslan/pgn/data/counter-vs-zahak.pgn")
}

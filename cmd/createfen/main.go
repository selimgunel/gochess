package main

import (
	_ "net/http/pprof"
)

func main() {

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	// // set up engine to use stockfish exe

	// // have stockfish play speed chess against itself (10 msec per move)
	// game := gochess.NewGame()

	// f, err := os.ReadFile("pgn/counter-vs-zahak.pgn")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// moves, _ := gochess.Parse(string(f))

	// for i, move := range moves {
	// 	fmt.Printf("[%d] %s fen: %s\n", i, move, game.FEN())
	// 	if err := game.MoveStr(move); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	//game.Move()
	// }

	// //gmoves := game.ValidMoves()
	// //game.Move(gmoves[0])
	// fmt.Println(game.FEN())

	// fmt.Println(game.Position().Board().Draw())

	// Output:

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/freeeve/pgn.v1"
)

const (
	Start = iota
	NewLine
	InTags
	InMoves
)

type Game struct {
	Tags   [][]string
	Moves  []string
	Result string
}

type Parser struct {
	Games []*Game
}

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// An artificial input source.
	f, err := os.Open("pgn/famous_games.pgn")
	//f, err := os.Open("data/counter-vs-zahak.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// splitPoint := make([]int, 0)
	// splitPoint = append(splitPoint, 0)
	// for i, v := range scanPGNBuf(f) {
	// 	if i%2 != 0 {
	// 		splitPoint = append(splitPoint, v)
	// 	}

	// }
	// f.Seek(0, 0)
	// games := extractGames(splitPoint, f)

	// //fmt.Println(games[0])
	// g := games[0]
	// g.ParseMoves()

	scanner := bufio.NewScanner(f)

	scanner.Split(crunchSplitFunc)

	sa := make([]string, 0)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		sa = append(sa, "["+t)
	}

	for i, v := range sa {
		fmt.Printf("[%d] %v\n", i, v)
	}

	fmt.Printf("length %d", len(sa))

	//FreevePGN()
}

func FreevePGN() {

	f, err := os.Open("pgn/famous_games.pgn")
	if err != nil {
		panic(err)
	}

	ps := pgn.NewPGNScanner(f)
	// while there's more to read in the file
	for ps.Next() {
		// scan the next game
		game, err := ps.Scan()
		if err != nil {
			log.Fatal().Msgf("%+v", err)
		}

		// print out tags
		fmt.Println(game.Tags)

		// make a new board so we can get FEN positions
		b := pgn.NewBoard()
		for _, move := range game.Moves {
			// make the move on the board
			b.MakeMove(move)
			// print out FEN for each move in the game
			fmt.Println(b)
		}
	}
}

//
func (g Game) String() string {
	var sb strings.Builder
	for _, v := range g.Tags {
		fmt.Fprintf(&sb, "%13s", v[0])
		sb.WriteString(" => ")
		fmt.Fprintf(&sb, "%s", v[1])
		sb.WriteRune('\n')
	}

	for i, v := range g.Moves {
		fmt.Fprintf(&sb, "%d: %s", i, v)
		sb.WriteRune('\n')
	}
	return sb.String()
}

// func scanPGNBuf(f *os.File) (cutPoints []int) {

// 	scanner := bufio.NewScanner(f)

// 	var lineNumber = 1

// 	cutPoints = make([]int, 0)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		if line == "" {
// 			cutPoints = append(cutPoints, lineNumber)
// 		}

// 		lineNumber++
// 		//log.Debug().Msg(line)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Printf("Invalid input: %s", err)
// 	}
// 	return cutPoints
// }

// func extractGames(splitPoints []int, f *os.File) []Game {

// 	scanner := bufio.NewScanner(f)
// 	stacks := make([]Game, 0)
// 	//var lineNumber = 1

// 	for j, ln := range splitPoints {
// 		//log.Debug().Msgf("ln %d", ln)
// 		var g Game
// 		if j == len(splitPoints)-1 {
// 			break
// 		}
// 		for i := ln; i < splitPoints[j+1]; i++ {

// 			scanner.Scan()

// 			//log.Debug().Msgf("%+v", scanner.Text())
// 			line := scanner.Text()
// 			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
// 				before, after, ok := strings.Cut(line, " \"")
// 				if ok {
// 					s := make([]string, 0)
// 					s = append(s, before[1:], after[:len(after)-2])
// 					g.Tags = append(g.Tags, s)
// 				}

// 			} else {
// 				if line != "" {
// 					g.Moves = append(g.Moves, line)
// 				}

// 			}

// 		}
// 		stacks = append(stacks, g)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal().Msgf("%+v", err)
// 	}
// 	return stacks
// }

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

//ParseMoves
func (g Game) ParseMoves() {

}

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "[Event"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	splitPoint := make([]int, 0)
	splitPoint = append(splitPoint, 0)
	for i, v := range scanPGNBuf(f) {
		if i%2 != 0 {
			splitPoint = append(splitPoint, v)
		}

	}
	f.Seek(0, 0)
	games := extractGames(splitPoint, f)

	fmt.Println(games[0])

	f.Close()
}

//
func (g Game) String() string {
	var sb strings.Builder
	for _, v := range g.Tags {
		fmt.Fprintf(&sb, "%10s", v[0])
		sb.WriteString(" => ")
		fmt.Fprintf(&sb, "%s", v[1])
		sb.WriteRune('\n')
	}
	return sb.String()
}

func scanPGNBuf(f *os.File) (cutPoints []int) {

	scanner := bufio.NewScanner(f)

	var lineNumber = 1

	cutPoints = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			cutPoints = append(cutPoints, lineNumber)
		}

		lineNumber++
		//log.Debug().Msg(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	return cutPoints
}

func extractGames(splitPoints []int, f *os.File) []Game {

	scanner := bufio.NewScanner(f)
	stacks := make([]Game, 0)
	//var lineNumber = 1

	for j, ln := range splitPoints {
		//log.Debug().Msgf("ln %d", ln)
		var g Game
		if j == len(splitPoints)-1 {
			break
		}
		for i := ln; i < splitPoints[j+1]; i++ {

			scanner.Scan()

			//log.Debug().Msgf("%+v", scanner.Text())
			line := scanner.Text()
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				before, after, ok := strings.Cut(line, " \"")
				if ok {
					s := make([]string, 0)
					s = append(s, before[1:], after[:len(after)-2])
					g.Tags = append(g.Tags, s)
				}

			} else {
				g.Moves = append(g.Moves, line)
			}

		}
		stacks = append(stacks, g)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal().Msgf("%+v", err)
	}
	return stacks
}

// IsEmpty: check if stack is empty
func (g *Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

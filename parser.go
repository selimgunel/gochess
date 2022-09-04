package parser

import (
	"bufio"
	"io"
	"strings"
)

const (
	Start = iota
	NewLine
	InTags
	InMoves
)

type Tag [2]string
type Game struct {
	Tags  []Tag
	Moves []string
}

func Parse(input io.Reader) []Game {

	scanner := bufio.NewScanner(input)
	scanner.Split(crunchSplitFunc)

	sa := make([]string, 0)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		sa = append(sa, "["+t)
	}

	games := make([]Game, 0)
	for _, src := range sa {
		l := NewLexer(strings.NewReader(src))

		g := Game{
			Tags:  make([]Tag, 0),
			Moves: make([]string, 0),
		}
		cc := 0
		for {

			tok := l.Scan()

			if tok.Name == EOF || tok.Name == ERROR {
				break
			}
			if tok.Name == MOVE && tok.Val == "" {
				continue
			}
			if tok.Name == COMMENT {
				continue
			}
			if tok.Name == NEWLINE {
				continue
			}

			if tok.Name == MOVE && tok.Val != "" {

				g.Moves = append(g.Moves, tok.Val)
			}
			cc++
		}

		games = append(games, g)
	}

	return games
}

//Tags

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

// Taken from: https://github.com/freeeve/pgn/issues/17.
func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "[Event "); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

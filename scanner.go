package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/narslan/schach/pgnparser/lexer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Start = iota
	NewLine
	InTags
	InMoves
)

type Tag [2]string
type Game struct {
	Tags   []Tag
	Moves  []string
	Result string
}

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// An artificial input source.
	//f, err := os.Open("pgn/famous_games.pgn")
	f, err := os.Open("data/counter-vs-zahak.pgn")
	//f, err := os.Open("data/anders.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()

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

	g := LexGameInput(sa[0])
	//fmt.Printf("%s", g)
	for _, v := range g.Tags {
		fmt.Printf("%s", v)
	}
}

func LexGameInput(src string) (g Game) {

	s := bufio.NewScanner(strings.NewReader(src))

	g.Tags = make([]Tag, 0)

	tags := make([]Tag, 0)
	var movesSrc strings.Builder
	for s.Scan() {

		l := s.Text()
		if l == "" {
			continue
		}

		if strings.Index(l, "[") == 0 {
			// splitted tags into key and value.
			//fmt.Printf("tag: %s", l)
			stags := strings.Split(l[1:len(l)-1], " ")
			k := stags[0]
			v := strings.ReplaceAll(stags[1], "\"", "")
			tagnames := [2]string{k, v}
			tags = append(tags, tagnames)
			continue
		}
		fmt.Fprintf(&movesSrc, "%s", l)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	parseMoves(movesSrc.String())
	g.Tags = tags
	return
}

func parseMoves(src string) {
	l := lexer.BeginLexing("deneme", src)

	//fmt.Printf("%s", src)
	for {
		token := l.NextToken()

		fmt.Printf("%+v\n", token)
	}
}

//Tags

func (t Tag) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%13s", t[0])
	sb.WriteString(" => ")
	fmt.Fprintf(&sb, "%s", t[1])
	sb.WriteRune('\n')

	return sb.String()
}

func (g Game) String() string {
	var sb strings.Builder
	for _, t := range g.Tags {
		fmt.Fprintf(&sb, "%s", t)
	}

	for i, v := range g.Moves {
		fmt.Fprintf(&sb, "%d: %s", i, v)
		sb.WriteRune('\n')
	}
	return sb.String()
}

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

//ParseMoves
func (g Game) ParseMoves() {

}

// Taken from: https://github.com/freeeve/pgn/issues/17.
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

package main

import (
	"fmt"
	"os"
	"strings"

	"net/http"
	_ "net/http/pprof"

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
	Tags   []*Tag
	Moves  string
	Result string
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// An artificial input source.
	//f, err := os.Open("pgn/famous_games.pgn")
	f, err := os.Open("/home/nevroz/go/src/github.com/narslan/schach/pgnparser/data/counter-vs-zahak.pgn")
	//f, err := os.Open("data/anders.pgn")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	l := lexer.NewLexer(f)

	for {

		tok := l.Scan()
		if tok.Name == lexer.EOF || tok.Name == lexer.ERROR {
			fmt.Printf("tok: %d %s Pos: %d\n", tok.Name, tok.Val, tok.Pos)
			break
		}
		fmt.Printf("tok: %d %s Pos: %d\n", tok.Name, tok.Val, tok.Pos)

	}

	// scanner := bufio.NewScanner(f)
	// scanner.Split(crunchSplitFunc)

	// sa := make([]string, 0)

	// for scanner.Scan() {
	// 	t := scanner.Text()
	// 	if t == "" {
	// 		continue
	// 	}
	// 	sa = append(sa, "["+t)
	// }

	//fmt.Printf("%d", len(sa))
	//	for _, v := range sa {
	//		fmt.Printf("%s\n", v)
	//	}
	//LexGameInput(sa[0])

	//fmt.Printf("%s", g)
	//g.ParseMoves()
	// for _, v := range g.Moves {
	// 	fmt.Printf("%s", v)
	// }
}

//Tags

func (t *Tag) String() string {
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

	//fmt.Fprintf(&sb, "%s", g.Moves)

	return sb.String()
}

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

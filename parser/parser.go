package parser

import (
	"bufio"
	"fmt"
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
	Tags   []*Tag
	Moves  string
	Result string
}

func Parse(input io.Reader) {

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

	for _, g := range sa {
		l := NewLexer(strings.NewReader(g))

		for {

			tok := l.Scan()
			if tok.Name == EOF || tok.Name == ERROR {
				fmt.Printf("tok: %d %s Pos: %d\n", tok.Name, tok.Val, tok.Pos)
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
			fmt.Printf("%s\n", tok)

		}
	}

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

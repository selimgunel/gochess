package pgn

import (
	"bufio"
	"fmt"
	"io"
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

func Parse(input io.Reader) ([]Game, error) {

	g := make([]Game, 0)
	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1
	for s.Scan() {
		l := s.Text()
		if l == "" {
			fmt.Printf("line number: %d\n", ln)
		}
		ln++
	}
	return g, nil
}

//Tags

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

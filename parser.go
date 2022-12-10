package pgn

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

func Split(input io.Reader) ([]Tag, []string, error) {

	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1

	tags := make([]Tag, 0)
	moves := make([]string, 0)
	var moveCtx bool //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
				tag := Tag{}
				i := strings.Index(l, " ")
				tag[0] = l[1:i]
				tag[1] = l[i:]
				tags = append(tags, tag)
			} else {
				moveCtx = true
				moves = append(moves, l)
			}

		} else {
			if moveCtx {
				tags = append(tags, Tag{})
				moves = append(moves, "---")
			}
		}
		ln++
	}

	return tags, moves, nil
}

func Parse(tags []Tag, moves []string) (*Game, error) {

	src := strings.Join(moves, "\n")
	l := NewLexer(strings.NewReader(src))

	g := &Game{
		Tags:  make([]Tag, 0),
		Moves: make([]string, 0),
	}
	g.Tags = tags
L:
	for {

		tok := l.Scan()

		switch tok.Name {
		case EOF, ERROR:
			break L
		case MOVE:
			if tok.Val != "" {
				g.Moves = append(g.Moves, tok.Val)
			}
		default:
			continue
		}

	}

	return g, nil
}

//Tags

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

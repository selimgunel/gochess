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
	var moveCtx = false //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
				tag := Tag{}
				i := strings.Index(l, " ")
				tag[0] = l[1:i]
				tag[1] = l[i : len(l)-1]
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

func SplitPoints(input io.Reader) ([]int, error) {

	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1

	splitPoints := make([]int, 0)
	var moveCtx = false //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
			} else {
				moveCtx = true

			}

		} else {
			if moveCtx {
				splitPoints = append(splitPoints, ln)
			}
		}

		ln++
	}

	return splitPoints, nil
}

func Parse(input io.Reader, sps []int) (*Game, error) {

	s := bufio.NewScanner(input)

	s.Split(bufio.ScanLines)
	ln := 1

	splitPoints := make([]int, 0)
	var moveCtx = false //whether the scanner is in tags or in moves
	for s.Scan() {
		l := s.Text()

		if l != "" {

			if strings.HasPrefix(l, "[") {
				moveCtx = false
			} else {
				moveCtx = true

			}

		} else {
			if moveCtx {
				splitPoints = append(splitPoints, ln)
			}
		}

		ln++
	}

}

//Tags

// IsEmpty: check if stack is empty
func (g Game) IsEmpty() bool {
	return len(g.Tags) == 0
}

`pgn` is a PGN file parser written in `go`. 
There is just one useful method, namely `gochessParse()`. 
```go
package main

import (
	parser "github.com/narslan/gochess"
	
)

func main() {

	//move ...
    file, err := os.Open("/path/to/pgnfile")
    //...
	games := parser.PosFromStart("PGN_FILE")

    for g := range games {
        //...    
    }
}
```

```go

func Parse(input io.Reader) ([]Game, error) {

	scanner := bufio.NewScanner(input)
	

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

			switch tok.Name {
			case EOF, ERROR:
				return games, nil
			case MOVE:
				if tok.Val != "" {
					g.Moves = append(g.Moves, tok.Val)
				}

			default:
				continue
			}

			cc++
		}

		games = append(games, g)
	}

	return games, nil
}```

```go

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

	
	for _, v := range tags {
		
	}

	return tags, moves, nil
}
```
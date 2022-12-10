`pgn` is a PGN file parser written in `go`. 
There is just one useful method, namely `pgn.Parse()`. 
```go
package main

import (
	parser "github.com/narslan/pgn"
	
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
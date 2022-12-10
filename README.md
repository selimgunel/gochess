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


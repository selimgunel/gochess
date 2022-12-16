package pgn

// Moves prints out moves.
func Parse(pgnsource string) (moves []string, tags []string) {

	lex := NewLexer(pgnsource)
	var toks []Token

	for {
		tok := lex.NextToken()

		toks = append(toks, tok)
		if tok.Name == EOF {
			break
		}

	}

	for _, v := range toks {
		if v.Name == MOVE {
			moves = append(moves, v.Val)
		}

		if v.Name == TAG {
			tags = append(tags, v.Val)
		}
	}
	return moves, tags
}

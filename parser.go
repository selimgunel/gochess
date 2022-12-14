package pgn

// Moves prints out moves.
func Moves(pgnsource string) []string {

	lex := NewLexer(pgnsource)
	var toks []Token

	for {
		tok := lex.NextToken()

		toks = append(toks, tok)
		if tok.Name == EOF {
			break
		}

	}
	tokens := make([]string, 0)
	for _, v := range toks {
		if v.Name == MOVE {
			tokens = append(tokens, v.Val)
		}

	}
	return tokens
}

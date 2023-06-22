package core

// Lex accepts raw source code as input and returns a slice of tokens.
func Lex(src string) []string {
	var tokens []string
	for i := 0; i < len(src); i++ {
		if token, ok := _tokens[src[i]]; ok {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || b == '_'
}

package lang

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	expr := "(begin (define r 10) (* pi (* r r)))"
	tokens := tokenize(expr)
	expressionize(tokens)
	t.Error()
}

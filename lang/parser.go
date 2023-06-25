package lang

import (
	"fmt"
	"strconv"
	"strings"
)

// Tokenize converts a string of characters into a series of tokens.
// Excess whitespace is removed automatically.
func tokenize(s string) (tokens []string) {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)
	for _, token := range strings.Split(s, " ") {
		if token == "" {
			continue
		}
		tokens = append(tokens, token)
	}
	return
}

// Atomize converts a string token an atom.
func atomize(token string) atom {
	// First, try converting the token into an integer.
	i, err := strconv.ParseInt(token, 0, 64)
	if err == nil {
		return atom{number, i}
	}

	// That failed, so try converting it to a float.
	f, err := strconv.ParseFloat(token, 64)
	if err == nil {
		return atom{number, f}
	}

	// Anything not a number is automatically a symbol, which
	// in this type system is represented by a string.
	return atom{symbol, token}
}

// Forward definitions for recursive etc. functions.
var (
	expand func() []atom
	pop    func(ts []string) (string, []string)
)

// Expressionize accepts a series of tokens and builds the appropriate
// expressions from them, which may be nested to an arbitrary depth.
func expressionize(tokens []string) (expr []atom) {
	var token string

	pop = func(ts []string) (string, []string) {
		return ts[0], ts[1:]
	}

	expand = func() []atom {
		token, tokens = pop(tokens)
		if token == "(" {
			var L []atom
			for tokens[0] != ")" {
				L = append(L, expand()...)
			}
			_, tokens = pop(tokens)
			return L
		} else if token == ")" {
			panic("unexpected )")
		} else {
			expr = append(expr, atomize(token))
			return nil
		}
	}

	expand()
	fmt.Println(expr)
	return
}

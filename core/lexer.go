package core

import (
	"unicode"
)

// Lex extracts tokens from the given golisp source.
func Lex(src string) (tokens Tokens) {
	// i holds the current position in src.
	// c holds the current byte in src.
	var i int
	var c byte

	var appendtok = func(id string, value string) {
		tokens = append(tokens, Token{ID: id, Value: value})
	}

	var isletter = func(value byte) bool {
		return unicode.IsLetter(rune(value)) || value == '_'
	}

	var isdigit = func(value byte) bool {
		return unicode.IsDigit(rune(value))
	}

	var iswhitespace = func(value byte) bool {
		return unicode.IsSpace(rune(value))
	}

	var readidentifier = func() {
		j := i
		for isletter(src[j]) {
			j++
		}
		identifier := src[i:j]
		appendtok("IDENTIFIER", identifier)
		i += len(identifier) - 1
	}

	var readnumber = func() {
		j := i
		for isdigit(src[j]) {
			j++
		}
		number := src[i:j]
		appendtok("NUMBER", number)
		i += len(number) - 1
	}

	for i = 0; i < len(src); i++ {
		c = src[i]
		switch c {
		case LPAREN:
			appendtok("LPAREN", string(c))
		case RPAREN:
			appendtok("RPAREN", string(c))
		case LBRACK:
			appendtok("LBRACK", string(c))
		case RBRACK:
			appendtok("RBRACK", string(c))
		case PLUS:
			appendtok("PLUS", string(c))
		default:
			if iswhitespace(c) {
				continue
			}
			if isletter(c) {
				readidentifier()
			} else if isdigit(c) {
				readnumber()
			} else {
				appendtok("ILLEGAL", string(c))
			}
		}
	}

	return
}

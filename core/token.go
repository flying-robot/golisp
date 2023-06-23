package core

// These constants define the possible language tokens.
const (
	// Delimiters
	LPAREN = '('
	RPAREN = ')'
	LBRACK = '['
	RBRACK = ']'
	SPACE  = ' '

	// Operators
	PLUS = '+'

	// System
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	ILLEGAL    = "ILLEGAL"
)

// A Token is the smallest subdivision of source code.
type Token struct {
	ID    string
	Value string
}

// Tokens contains zero or more tokens, extracted via lexing.
type Tokens []Token

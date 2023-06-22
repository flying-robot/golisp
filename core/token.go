package core

// This map defines the tokens known to the language.
var _tokens = map[byte]string{
	'(': "LPAREN",
	')': "RPAREN",
	'{': "LBRACK",
	'}': "RBRACK",
}

// This map defines the keywords known to the language.
var _keywords = map[string]string{
	"def":  "DEF",
	"defn": "DEFN",
}

package core

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		desc string
		src  string
		res  Tokens
	}{
		{
			desc: "A simple program with parens",
			src:  "()",
			res: Tokens{
				{"LPAREN", "("},
				{"RPAREN", ")"},
			},
		},
		{
			desc: "A simple program with parens and an identifier",
			src:  "(hello)",
			res: Tokens{
				{"LPAREN", "("},
				{"IDENTIFIER", "hello"},
				{"RPAREN", ")"},
			},
		},
		{
			desc: "A slightly more complex program with nested parens",
			src:  "(+ (+ 1 2) 3)",
			res: Tokens{
				{"LPAREN", "("},
				{"PLUS", "+"},
				{"LPAREN", "("},
				{"PLUS", "+"},
				{"NUMBER", "1"},
				{"NUMBER", "2"},
				{"RPAREN", ")"},
				{"NUMBER", "3"},
				{"RPAREN", ")"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := Lex(test.src)
			if !reflect.DeepEqual(got, test.res) {
				t.Fatalf("got %v, want %v", got, test.res)
			}
		})
	}
}

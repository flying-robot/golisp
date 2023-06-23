// Package main provides the entry point to the golisp repl.
package main

import (
	"fmt"
	"golisp/core"

	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("golisp v0.0.1")
	fmt.Println("Press CTRL+C to exit the REPL session.")
	fmt.Println("")
	repl()
}

func repl() {
	rl, err := readline.New("golisp> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		for _, token := range core.Lex(line) {
			fmt.Println(token)
		}
		fmt.Println()
	}
}

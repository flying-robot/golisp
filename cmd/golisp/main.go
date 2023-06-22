// Package main provides the entry point to the golisp repl.
package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

func main() {
	fmt.Println("golisp v0.0.1")
	fmt.Println("Press CTRL+C to exit the REPL session.")
	fmt.Println("")
	acceptInput()
}

func acceptInput() {
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
		println(line)
	}
}

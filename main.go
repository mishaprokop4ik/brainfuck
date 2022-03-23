package main

import (
	"github.com/mishaprokop4ik/brainfuck/internal/interpreter"
)

// nolint
const defaultDataIn string = `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.`

func main() {
	interpreter.Interpret(defaultDataIn)
}

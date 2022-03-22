package main

import (
	"github.com/mishaprokop4ik/brainfuck/internal/interpreter"
)

const defaultDataIn string = ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.\n"

func main()  {
	interpreter.Interpret(defaultDataIn)
}

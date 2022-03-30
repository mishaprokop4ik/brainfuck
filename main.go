package main

import (
	"github.com/mishaprokop4ik/brainfuck/internal/interpreter"
)

//const defaultDataIn string = ">++++++++[<+++++++++>-]" +
//	"<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++" +
//	"[<+++++++>-]<++.------------.>++++++" +
//	"[<+++++++++>-]<+.<.+++.------.--------" +
//	".>>>++++[<++++++++>-]<+."

func main() {
	interpreter.Interpret("-[------->+<]>-.-[->+++++<]>++.+++++++..+++.", nil)
}

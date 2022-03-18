package main

import (
	"github.com/mishaprokop4ik/brainfuck/internal/interpreter"
)

func main()  {
	input := "-[--->+<]>-.[---->+++++<]>-.+.++++++++++.+[---->+<]>+++.-[--->++<]>-.++++++++++.+[---->+<]>+++.[-->+++++++<]>.++.-------------.[--->+<]>---..+++++.-[---->+<]>++.+[->+++<]>.++++++++++++..---.[-->+<]>--------."
	interpreter.Interpret(input)
}

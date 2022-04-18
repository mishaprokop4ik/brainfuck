package interpreter

import (
	"fmt"
)

// Interpret gets a brainfuck operations as a string.
// These operations will be converted into go methods
// and after it all methods will be run
// if some unexpected symbol in input - program will panic
func Interpret(program string) {
	operations := findOperations(program)
	m := newMemory(memorySize)

	for _, o := range operations {
		o.execute(m)
	}
}

func findOperations(program string) []operation {
	var stack = &[][]operation{{}}
	for i, symbol := range program {
		if instruction, ok := instructions[symbol]; ok {
			instruction(stack)
		} else {
			panic(fmt.Sprintf("unexpected symbol in brainfuck %s in %d place", string(symbol), i+1))
		}
	}

	return (*stack)[len(*stack)-1]
}

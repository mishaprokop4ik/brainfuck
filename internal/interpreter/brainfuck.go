package interpreter

import (
	"fmt"
	"io"
	"os"
)

// Interpret gets a brainfuck operations as a string.
// These operations will be converted into go methods
// and after it all methods will be run
// if some unexpected symbol in input - will be panic
func Interpret(program string, out *io.Writer) {
	operations := findOperations(program)
	var m *memory
	if out != nil {
		m = newMemory(memorySize, *out)
	} else {
		m = newMemory(memorySize, os.Stdout)
	}

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

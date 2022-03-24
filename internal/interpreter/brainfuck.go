package interpreter

import "fmt"

func Interpret(program string) {
	operations := findCommands(program)

	m := newMemory(memorySize)

	for _, o := range operations {
		o.execute(m)
	}
}

func findCommands(program string) []operation {
	var stack = &[][]operation{{}}
	for i, symbol := range program {
		if inst, ok := instructions[symbol]; ok {
			inst(stack)
		} else {
			panic(fmt.Sprintf("unexpected symbol in brainfuck %s in %d place", string(symbol), i+1))
		}
	}

	return (*stack)[len(*stack)-1]
}

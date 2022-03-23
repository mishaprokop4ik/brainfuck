package interpreter

import "fmt"

func Interpret(program string) {
	operations := findCommands(program)

	m := newMemory(memorySize)

	for _, o := range operations {
		o.Execute(m)
	}
	fmt.Println()
}

func findCommands(program string) []operation {
	var stack = &[][]operation{{}}
	for i, symbol := range program {
		if o, ok := commands[symbol]; ok {
			o(stack)
		} else {
			panic(fmt.Sprintf("unexpected symbol in brainfuck %s in %d place", symbol, i+1))
		}
	}

	return (*stack)[len(*stack)-1]
}

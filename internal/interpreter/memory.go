package interpreter

import "io"

const memorySize int = 30000

type memory struct {
	cells   []byte
	pointer int
	out     io.Writer
}

func newMemory(size int, output io.Writer) *memory {
	return &memory{
		cells:   make([]byte, size),
		pointer: 0,
		out:     output,
	}
}

func addOperation(stack *[][]operation, o operation) {
	operations := (*stack)[len(*stack)-1]
	(*stack)[len(*stack)-1] = append(operations, o)
}

type instruction func(stack *[][]operation)

var instructions = map[rune]instruction{
	'>': func(stack *[][]operation) {
		addOperation(stack, incrementPointer{})
	},
	'<': func(stack *[][]operation) {
		addOperation(stack, decrementPointer{})
	},
	'+': func(stack *[][]operation) {
		addOperation(stack, incrementCell{})
	},
	'-': func(stack *[][]operation) {
		addOperation(stack, decrementCell{})
	},
	'.': func(stack *[][]operation) {
		addOperation(stack, output{})
	},
	'[': func(stack *[][]operation) {
		*stack = append(*stack, []operation{})
	},
	']': func(stack *[][]operation) {
		loopStack := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		addOperation(stack, &loop{innerOperations: loopStack})
	},
}

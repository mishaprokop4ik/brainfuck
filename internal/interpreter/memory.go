package interpreter

const memorySize int = 30000

type memory struct {
	cells []byte
	pointer int
}

func newMemory(m int) *memory {
	return &memory{
		cells:   make([]byte, m),
		pointer: 0,
	}
}

func addOperation(stack *[][]operation, o operation) {
	operations := (*stack)[len(*stack) - 1]
	(*stack)[len(*stack) - 1] = append(operations, o)
}

type instruction func(stack *[][]operation)

var commands = map[rune]instruction{
	'>': func(stack *[][]operation) {
		addOperation(stack, IncrementPointer{})
	},
	'<': func(stack *[][]operation) {
		addOperation(stack, DecrementPointer{})
	},
	'+': func(stack *[][]operation) {
		addOperation(stack, IncrementCell{})
	},
	'-': func(stack *[][]operation) {
		addOperation(stack, DecrementCell{})
	},
	'.': func(stack *[][]operation) {
		addOperation(stack, Output{})
	},
	'[': func(stack *[][]operation) {
		*stack = append(*stack, []operation{})
	},
	']': func(stack *[][]operation) {
		loopStack := (*stack)[len(*stack) - 1]
		*stack = (*stack)[:len(*stack) - 1]
		addOperation(stack, &Loop{innerOperations: loopStack})
	},
}
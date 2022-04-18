package interpreter

import "fmt"

type operation interface {
	execute(m *memory)
}

// incrementPointer adds to pointer one
// if the pointer looks at the last cell
// pointer moves to the start
type incrementPointer struct{}

func (incrementPointer) execute(m *memory) {
	if m.pointer == len(m.cells)-1 {
		m.pointer = 0
	} else {
		m.pointer++
	}
}

// decrementPointer minuses pointer by one
// if the pointer looks at the first cell
// pointer moves to the end
type decrementPointer struct{}

func (decrementPointer) execute(m *memory) {
	if m.pointer == 0 {
		m.pointer = len(m.cells) - 1
	} else {
		m.pointer--
	}
}

// incrementCell adds to the current cell one
type incrementCell struct{}

func (incrementCell) execute(m *memory) {
	m.cells[m.pointer]++
}

// decrementCell minuses one from the current cell
type decrementCell struct{}

func (decrementCell) execute(m *memory) {
	m.cells[m.pointer]--
}

// output prints current cell to os.Stdout in ASCII representation
type output struct{}

func (output) execute(m *memory) {
	fmt.Print(string(m.cells[m.pointer]))
}

// loop contains inner operations and runs them while cell
// to which pointer indexes is not equal to zero
type loop struct {
	innerOperations []operation
}

func (l *loop) execute(m *memory) {
	for m.cells[m.pointer] != 0 {
		for _, o := range l.innerOperations {
			o.execute(m)
		}
	}
}

package interpreter

import "fmt"

type operation interface {
	execute(m *memory)
}

// incrementPointer it is a struct that implement operation interface
// it increment pointer by one
// if pointer points to the last cell
// pointer moves to the start
type incrementPointer struct{}

func (incrementPointer) execute(m *memory) {
	if m.pointer == len(m.cells)-1 {
		m.pointer = 0
	} else {
		m.pointer++
	}
}

// decrementPointer it is a struct that implement operation interface
// it increment pointer by one
// if pointer points to the first cell
// pointer moves to the end
type decrementPointer struct{}

func (decrementPointer) execute(m *memory) {
	if m.pointer == 0 {
		m.pointer = len(m.cells) - 1
	} else {
		m.pointer--
	}
}

// incrementCell it is a struct that implement operation interface
// add to the current cell one
type incrementCell struct{}

func (incrementCell) execute(m *memory) {
	m.cells[m.pointer]++
}

// decrementCell it is a struct that implement operation interface
// minus to the current cell one
type decrementCell struct{}

func (decrementCell) execute(m *memory) {
	m.cells[m.pointer]--
}

// output it is a struct that implement operation interface
// output current cell to stdout from byte to ASCII symbol
type output struct{}

func (output) execute(m *memory) {
	fmt.Print(string(m.cells[m.pointer]))
}

// loop it is a struct that implement operation interface
// operation in loop will be reproduced while current cell != zero
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

package interpreter

import "fmt"

type operation interface {
	Execute(m *memory)
}

type IncrementPointer struct {}

func (IncrementPointer) Execute(m *memory) {
	if m.pointer == len(m.cells) - 1 {
		m.pointer = 0
	} else {
		m.pointer++
	}
}

type DecrementPointer struct {}

func (DecrementPointer) Execute(m *memory) {
	if m.pointer == 0 {
		m.pointer = len(m.cells) - 1
	} else {
		m.pointer--
	}
}

type IncrementCell struct {}

func (IncrementCell) Execute(m *memory) {
	m.cells[m.pointer]++
}

type DecrementCell struct {}

func (DecrementCell) Execute(m *memory) {
	m.cells[m.pointer]--
}

type Output struct {}

func (Output) Execute(m *memory) {
	fmt.Print(string(m.cells[m.pointer]))
}

type Loop struct {
	innerOperations []operation
}

func (l *Loop) Execute(m *memory) {
	for m.cells[m.pointer] != 0 {
		for _, o := range l.innerOperations {
			o.Execute(m)
		}
	}
}

package interpreter

import "fmt"

type operation interface {
	execute(m *memory)
}

type IncrementPointer struct{}

func (IncrementPointer) execute(m *memory) {
	if m.pointer == len(m.cells)-1 {
		m.pointer = 0
	} else {
		m.pointer++
	}
}

type DecrementPointer struct{}

func (DecrementPointer) execute(m *memory) {
	if m.pointer == 0 {
		m.pointer = len(m.cells) - 1
	} else {
		m.pointer--
	}
}

type IncrementCell struct{}

func (IncrementCell) execute(m *memory) {
	m.cells[m.pointer]++
}

type DecrementCell struct{}

func (DecrementCell) execute(m *memory) {
	m.cells[m.pointer]--
}

type Output struct{}

func (Output) execute(m *memory) {
	fmt.Print(string(m.cells[m.pointer]))
}

type Loop struct {
	innerOperations []operation
}

func (l *Loop) execute(m *memory) {
	for m.cells[m.pointer] != 0 {
		for _, o := range l.innerOperations {
			o.execute(m)
		}
	}
}

package interpreter

import (
	"reflect"
	"testing"
)

func TestIncrementCell_Execute(t *testing.T) {
	tests := []struct {
		name      string
		memory    *memory
		expected  memory
		operation operation
	}{
		{
			name: "increment cell",
			memory: &memory{
				cells:   []byte{0, 0},
				pointer: 0,
			},
			expected: memory{
				cells:   []byte{1, 0},
				pointer: 0,
			},
			operation: IncrementCell{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.Execute(tt.memory)
			res := tt.memory
			if !reflect.DeepEqual(tt.expected, *res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

func TestDecrementCell_Execute(t *testing.T) {
	tests := []struct {
		name      string
		args      *memory
		expected  memory
		operation operation
	}{
		{
			name: "decrement cell",
			args: &memory{
				cells:   []byte{1, 0},
				pointer: 0,
			},
			expected: memory{
				cells:   []byte{0, 0},
				pointer: 0,
			},
			operation: DecrementCell{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.Execute(tt.args)
			res := tt.args
			if !reflect.DeepEqual(tt.expected, *res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

func TestIncrementPointer_Execute(t *testing.T) {
	tests := []struct {
		name      string
		memory    *memory
		expected  memory
		operation operation
	}{
		{
			name: "increment pointer",
			memory: &memory{
				cells:   []byte{},
				pointer: 0,
			},
			expected: memory{
				cells:   []byte{},
				pointer: 1,
			},
			operation: IncrementPointer{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.Execute(tt.memory)
			res := tt.memory
			if !reflect.DeepEqual(tt.expected, *res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

func TestDecrementPointer_Execute(t *testing.T) {
	tests := []struct {
		name      string
		memory    *memory
		expected  memory
		operation operation
	}{
		{
			name: "increment pointer",
			memory: &memory{
				cells:   []byte{},
				pointer: 1,
			},
			expected: memory{
				cells:   []byte{},
				pointer: 0,
			},
			operation: DecrementPointer{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.Execute(tt.memory)
			res := tt.memory
			if !reflect.DeepEqual(tt.expected, *res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

func TestLoop_Execute(t *testing.T) {
	tests := []struct {
		name          string
		args          *memory
		expected      memory
		preOperations []operation
		loopCommands  []operation
	}{
		{
			name: "loop execute with non zero value pointer",
			args: &memory{
				cells:   make([]byte, 5),
				pointer: 0,
			},
			preOperations: []operation{IncrementCell{}, IncrementCell{}, IncrementCell{}, IncrementCell{}, IncrementCell{}},
			expected: memory{
				cells:   make([]byte, 5),
				pointer: 0,
			},
			loopCommands: []operation{DecrementCell{}},
		},
		{
			name: "loop execute with zero value pointer",
			args: &memory{
				cells:   make([]byte, 5),
				pointer: 0,
			},
			expected: memory{
				cells:   make([]byte, 5),
				pointer: 0,
			},
			loopCommands: []operation{IncrementPointer{}, IncrementCell{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.preOperations != nil {
				for _, o := range tt.preOperations {
					o.Execute(tt.args)
				}
			}

			l := Loop{innerOperations: tt.loopCommands}

			l.Execute(tt.args)

			if !reflect.DeepEqual(tt.expected, *tt.args) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, tt.args)
			}
		})
	}
}

//nolint
package interpreter

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestOutput_Execute(t *testing.T) {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()
	tests := []struct {
		name      string
		memory    *memory
		operation operation
		expected  *bytes.Buffer
	}{
		{
			name: "output big A",
			memory: &memory{
				cells:   []byte{65},
				pointer: 0,
			},
			operation: output{},
			expected:  bytes.NewBuffer([]byte{65}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.execute(tt.memory)
			w.Close()
			res, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.expected.Bytes(), res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

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
			operation: incrementCell{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation.execute(tt.memory)
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
			operation: decrementCell{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint scopelint
			tt.operation.execute(tt.args)
			// nolint scopelint
			res := tt.args
			// nolint scopelint
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
			operation: incrementPointer{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint scopelint
			tt.operation.execute(tt.memory)
			// nolint scopelint
			res := tt.memory
			// nolint scopelint
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
			operation: decrementPointer{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// nolint scopelint
			tt.operation.execute(tt.memory)
			// nolint scopelint
			res := tt.memory
			// nolint scopelint
			if !reflect.DeepEqual(tt.expected, *res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, res)
			}
		})
	}
}

func TestLoop_Execute(t *testing.T) {
	tests := []struct {
		name         string
		args         *memory
		expected     memory
		loopCommands []operation
	}{
		{
			name: "loop execute with non zero value pointer",
			args: &memory{
				cells:   []byte{5},
				pointer: 0,
			},
			expected: memory{
				cells:   []byte{0},
				pointer: 0,
			},
			loopCommands: []operation{decrementCell{}},
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
			loopCommands: []operation{incrementPointer{}, incrementCell{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := loop{innerOperations: tt.loopCommands}
			l.execute(tt.args)
			if !reflect.DeepEqual(tt.expected, *tt.args) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected, tt.args)
			}
		})
	}
}

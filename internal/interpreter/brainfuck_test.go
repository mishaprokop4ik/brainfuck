package interpreter

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			"hello world program",
			"-[------->+<]>-.-[->+++++<]>++.+++++++..+++.[--->+<]>-----.--[->++++<]>-.--------.+++.------.--------.-[--->+<]>.",
		},
		{
			"unexpected symbol",
			"p",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() { recover() }()
			// nolint scopelint
			Interpret(tt.args)
		})
	}
}

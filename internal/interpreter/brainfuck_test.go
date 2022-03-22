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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interpret(tt.args)
		})
	}
}

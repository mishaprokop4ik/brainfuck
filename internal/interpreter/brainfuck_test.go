package interpreter

import (
	"fmt"
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
			defer func() {
				if p := recover(); p != nil && tt.name == "unexpected symbol" {
					fmt.Println(p)
				}
			}()
			// nolint scopelint
			Interpret(tt.args)
		})
	}
}

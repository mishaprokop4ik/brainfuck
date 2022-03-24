package interpreter

import (
	"fmt"
	"testing"
)

func ExampleInterpretHelloWorld() {
	Interpret("-[------->+<]>-.-[->+++++<]>++.+++++++..+++.[--->+<]>-----.--[->++++<]>-.--------.+++.------.--------.-[--->+<]>.")
	// Output:
	// Hello world!
}

func TestInterpret(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
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

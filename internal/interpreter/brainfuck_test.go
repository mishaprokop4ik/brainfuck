package interpreter

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		expected *bytes.Buffer
		out      io.Writer
	}{
		{
			"Hello program",
			"-[------->+<]>-.-[->+++++<]>++.+++++++..+++.",
			bytes.NewBuffer([]byte{72, 101, 108, 108, 111}),
			bytes.NewBuffer([]byte{}),
		},
		{
			"unexpected symbol",
			"p",
			bytes.NewBuffer([]byte{}),
			bytes.NewBuffer([]byte{}),
		},
		{
			"empty input",
			"",
			bytes.NewBuffer([]byte{}),
			bytes.NewBuffer([]byte{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				//nolint
				if p := recover(); p != nil && tt.name != "unexpected symbol" {
					t.Fatal(tt.name)
				}
			}()
			// nolint scopelint
			Interpret(tt.args, &tt.out)
			// nolint scopelint
			if v, ok := tt.out.(*bytes.Buffer); ok {
				// nolint scopelint
				if !reflect.DeepEqual(tt.expected.Bytes(), v.Bytes()) {
					// nolint scopelint
					t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected.Bytes(), v.Bytes())
				}
			}
		})
	}
}

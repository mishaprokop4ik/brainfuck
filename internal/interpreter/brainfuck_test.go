// nolint
package interpreter

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestInterpret(t *testing.T) {
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
		name     string
		args     string
		expected *bytes.Buffer
	}{
		{
			"hello program",
			"-[------->+<]>-.-[->+++++<]>++.+++++++..+++.",
			bytes.NewBuffer([]byte{72, 101, 108, 108, 111}),
		},
		{
			"unexpected symbol",
			"p",
			bytes.NewBuffer([]byte{}),
		},
		{
			"empty input",
			"",
			bytes.NewBuffer([]byte{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if p := recover(); p != nil && tt.name != "unexpected symbol" {
					t.Fatal(tt.name)
				}
			}()
			Interpret(tt.args)
			w.Close()
			res, err := ioutil.ReadAll(r)
			if err != nil && err != io.EOF {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.expected.Bytes(), res) {
				t.Errorf("Test failed %s expected: %v, got: %v", tt.name, tt.expected.Bytes(), res)
			}
		})
	}
}

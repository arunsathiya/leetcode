package main

import (
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		input  []byte
		output int
		result []byte
	}{
		{input: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, output: 6, result: []byte{'a', '2', 'b', '2', 'c', '3'}},
		{input: []byte{'a'}, output: 1, result: []byte{'a'}},
		{input: []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, output: 4, result: []byte{'a', 'b', '1', '2'}},
	}

	for _, test := range tests {
		if got := compress(test.input); got != test.output {
			t.Errorf("compress(%v) = %v; want %v", test.input, got, test.output)
		}
		for i := 0; i < test.output; i++ {
			if test.input[i] != test.result[i] {
				t.Errorf("compress(%v) modified input to %v; want %v", test.input, test.input, test.result)
				break
			}
		}
	}
}

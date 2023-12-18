package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "test case 2",
			input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			output: 9,
		},
		{
			name:   "test case 1",
			input:  []int{100, 4, 200, 1, 3, 2},
			output: 4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := longestConsecutive(test.input)
			if output != test.output {
				t.Errorf("Expected %d, got %d", test.output, output)
			}
		})
	}
}

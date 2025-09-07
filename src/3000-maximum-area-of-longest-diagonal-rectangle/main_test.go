package main

import (
	"testing"
)

func TestAreaOfMaxDiagonal(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name:     "first test case from leetcode",
			input:    [][]int{{9, 3}, {8, 6}},
			expected: 48,
		},
		{
			name:     "second test case from leetcode",
			input:    [][]int{{3, 4}, {4, 3}},
			expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := areaOfMaxDiagonal(test.input)
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}

func BenchmarkAreaOfMaxDiagonal(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// areaOfMaxDiagonal(...)
	}
}

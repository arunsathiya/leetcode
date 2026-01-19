package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "first test case from Leetcode",
			input:    123,
			expected: 321,
		},
		{
			name:     "second test case from leetcode",
			input:    -123,
			expected: -321,
		},
		{
			name:     "third test case from leetcode",
			input:    120,
			expected: 21,
		},
		{
			name:     "fourth test case from leetcode",
			input:    1534236469,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := reverse(tt.input)
			if output != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, output)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	// TODO: add benchmark inputs
	for i := 0; i < b.N; i++ {
		// reverse(...)
	}
}

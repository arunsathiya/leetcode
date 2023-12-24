package main

import (
	"testing"
)

func TestFindMin2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case from leetcode",
			input:    []int{1, 3, 5},
			expected: 1,
		},
		{
			name:     "second test case from leetcode",
			input:    []int{2, 2, 2, 0, 1},
			expected: 0,
		},
		{
			name:     "third test case from leetcode",
			input:    []int{10, 1, 10, 10, 10},
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := findMin2(test.input)
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}

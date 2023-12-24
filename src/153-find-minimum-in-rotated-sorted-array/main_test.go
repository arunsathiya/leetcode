package main

import "testing"

func TestFindMin1(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case from leetcode",
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "second test case from leetcode",
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "third test case from leetcode",
			input:    []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "fourth test case from leetcode",
			input:    []int{2, 1},
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := findMin1(test.input)
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}

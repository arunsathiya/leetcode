package main

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "first test case from leetcode",
			n:        2,
			expected: 1,
		},
		{
			name:     "second test case from leetcode",
			n:        3,
			expected: 2,
		},
		{
			name:     "third test case from leetcode",
			n:        4,
			expected: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := fib(test.n)
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}

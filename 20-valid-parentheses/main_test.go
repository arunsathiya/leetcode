package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "first test case from leetcode",
			input:    "()",
			expected: true,
		},
		{
			name:     "second test case from leetcode",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "third test case from leetcode",
			input:    "(]",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isValid(test.input)
			if output != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}

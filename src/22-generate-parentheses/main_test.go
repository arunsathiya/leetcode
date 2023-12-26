package main

import "testing"

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "first test case from leetcode",
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "second test case from leetcode",
			input:    1,
			expected: []string{"()"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := generateParenthesis(test.input)
			if !equalityWithoutOrder(output, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}

func equalityWithoutOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[string]int)
	for _, item := range a {
		count[item]++
	}
	for _, item := range b {
		if count[item] == 0 {
			return false
		}
		count[item]--
	}
	return true
}

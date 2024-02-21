package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // Use a slice of strings since there can be more than one valid answer
	}{
		{
			name:     "Example 1",
			input:    "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			expected: []string{"bb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestPalindrome(tt.input)
			assert.Contains(t, tt.expected, actual) // Assert that the actual answer is within the expected answers
		})
	}
}

package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Example 1", "abcabcbb", 3},
		{"Example 2", "bbbbb", 1},
		{"Example 3", "pwwkew", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tc.input)
			if actual != tc.expected {
				t.Errorf("For input '%s', expected length %d but got %d", tc.input, tc.expected, actual)
			}
		})
	}
}

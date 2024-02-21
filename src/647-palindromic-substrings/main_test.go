package main

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "abc",
			expected: 3, // Explanation: Three palindromic strings: "a", "b", "c".
		},
		{
			s:        "aaa",
			expected: 6, // Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			result := countSubstrings(tt.s)
			if result != tt.expected {
				t.Errorf("countSubstrings(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}

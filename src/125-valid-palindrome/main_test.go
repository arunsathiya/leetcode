package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "first test case from leetcode",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "second test case from leetcode",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "third test case from leetcode",
			s:        "",
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isPalindrome(test.s)
			if output != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}

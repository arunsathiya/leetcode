package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		s        string // input string
		expected bool   // expected result
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			result := isPalindrome(tc.s)
			if result != tc.expected {
				t.Errorf("For input '%s', expected %v but got %v", tc.s, tc.expected, result)
			}
		})
	}
}

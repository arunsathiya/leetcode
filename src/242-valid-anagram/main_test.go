package main

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},          // Edge case: both strings are empty
		{"a", "a", true},        // Edge case: both strings contain the same single character
		{"abc", "cba", true},    // Test case: simple anagram with no repeated letters
		{"test", "ttew", false}, // Test case: same length but not an anagram
	}

	for _, tc := range testCases {
		actual := isAnagram(tc.s, tc.t)
		if actual != tc.expected {
			t.Errorf("isAnagram(%v, %v) = %v, expected %v", tc.s, tc.t, actual, tc.expected)
		}
	}
}

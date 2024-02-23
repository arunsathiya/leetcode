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
	}

	for _, tc := range testCases {
		actual := isAnagram(tc.s, tc.t)
		if actual != tc.expected {
			t.Errorf("isAnagram(%v, %v) = %v, expected %v", tc.s, tc.t, actual, tc.expected)
		}
	}
}

package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, test := range tests {
		if result := isSubsequence(test.s, test.t); result != test.expected {
			t.Errorf("isSubsequence(%q, %q) = %v; want %v", test.s, test.t, result, test.expected)
		}
	}
}

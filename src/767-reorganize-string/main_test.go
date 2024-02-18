package main

import "testing"

func TestReorganizeString(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"aab", "aba"},
		{"aaab", ""},
	}

	for _, tc := range testCases {
		result := reorganizeString(tc.s)
		if result != tc.expected && (len(result) == len(tc.s) && len(result) > 0) {
			for i := 1; i < len(result); i++ {
				if result[i] == result[i-1] {
					t.Errorf("reorganizeString(%v) = %v, expected %v", tc.s, result, tc.expected)
					break
				}
			}
		} else if result != tc.expected {
			t.Errorf("reorganizeString(%v) = %v, expected %v", tc.s, result, tc.expected)
		}
	}
}

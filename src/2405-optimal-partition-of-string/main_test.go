package main

import (
	"testing"
)

func TestPartitionString(t *testing.T) {
	testCases := []struct {
		s        string
		expected int
	}{
		{"abacaba", 4},
		{"ssssss", 6},
	}

	for _, tc := range testCases {
		result := partitionString(tc.s)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %d but got %d", tc.s, tc.expected, result)
		}
	}
}

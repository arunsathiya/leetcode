package main

import "testing"

func TestKthFactor(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected int
	}{
		{12, 3, 3},
		{7, 2, 7},
		{4, 4, -1},
	}

	for _, tc := range testCases {
		result := kthFactor(tc.n, tc.k)
		if result != tc.expected {
			t.Errorf("For n = %d, k = %d, expected %d but got %d", tc.n, tc.k, tc.expected, result)
		}
	}
}

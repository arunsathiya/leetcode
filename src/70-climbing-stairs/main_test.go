package main

import "testing"

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"2 steps", 2, 2},
		{"3 steps", 3, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := climbStairs(tc.input)
			if result != tc.expected {
				t.Errorf("For input '%d', expected '%d' but got '%d'", tc.input, tc.expected, result)
			}
		})
	}
}

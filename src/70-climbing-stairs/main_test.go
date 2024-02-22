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
		{"1 step", 1, 1},  // Adding a test case for the base condition.
		{"4 steps", 4, 5}, // Adding a test case for 4 steps, expected outcome is 5 ways.
		{"5 steps", 5, 8}, // Adding a test case for 5 steps, expected outcome is 8 ways.
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

package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "first test case from leetcode",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "second test case from leetcode",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := search(test.nums, test.target)
			if output != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, output)
			}
		})
	}
}

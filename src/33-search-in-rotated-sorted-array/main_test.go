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
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "second test case from leetcode",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "third test case from leetcode",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
	}
	for _, test := range tests {
		output := search(test.nums, test.target)
		if output != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, output)
		}
	}
}

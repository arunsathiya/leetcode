package main

import "testing"

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1, 5, 0, 3, 5}, 3},
		{[]int{0}, 0},
	}

	for _, test := range tests {
		result := minimumOperations(test.nums)
		if result != test.output {
			t.Errorf("For nums=%v, expected output=%d, but got %d", test.nums, test.output, result)
		}
	}
}

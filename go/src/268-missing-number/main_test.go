package main

import "testing"

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tc := range testCases {
		got := missingNumber(tc.nums)
		if got != tc.want {
			t.Errorf("For input %v; expected %d but got %d", tc.nums, tc.want, got)
		}
	}
}

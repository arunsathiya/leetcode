package main

import "testing"

func TestCanJump(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1: Reachable last index",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Example 2: Unreachable last index",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "Large jump at start",
			nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 4, 5},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canJump(tc.nums)
			if got != tc.want {
				t.Errorf("canJump(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

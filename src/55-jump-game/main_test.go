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

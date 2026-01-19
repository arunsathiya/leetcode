package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "test1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "test2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "test3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

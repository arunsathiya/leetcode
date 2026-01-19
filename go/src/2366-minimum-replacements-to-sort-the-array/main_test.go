package main

import "testing"

func TestMinimumReplacement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{3, 9, 3},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "Example 3",
			nums: []int{12, 9, 7, 6, 17, 19, 21},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumReplacement(tt.nums); got != tt.want {
				t.Errorf("minimumReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
			want:  [][]int{{1, 3}, {4, 6}},
		},
		{
			name:  "Example 2",
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			want:  [][]int{{3}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

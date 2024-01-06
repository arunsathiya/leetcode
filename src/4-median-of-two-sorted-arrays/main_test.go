package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "first test case from leetcode",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.00000,
		},
		{
			name:     "second test case from leetcode",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.50000,
		},
		{
			name:     "third test case from leetcode",
			nums1:    []int{1, 3},
			nums2:    []int{2, 7},
			expected: 2.50000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := findMedianSortedArrays(test.nums1, test.nums2)
			if output != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, output)
			}
		})
	}
}

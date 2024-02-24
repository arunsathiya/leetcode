package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	}
	lower := (len(nums1) - 1) / 2
	upper := lower + 1
	return (float64(nums1[lower]) + float64(nums1[upper])) / 2
}

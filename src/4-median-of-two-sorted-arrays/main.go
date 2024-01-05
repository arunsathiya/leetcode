package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	sum := 0
	for _, num := range nums1 {
		sum += num
	}
	return float64(sum) / float64(len(nums1))
}

package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[m+i] = nums2[i]
		}
	}
	sort.Ints(nums1)
}

package main

import "sort"

func findMin(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})
	return nums[0]
}

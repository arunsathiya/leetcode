package main

import "sort"

func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	j := n - 1
	for i := 0; i < n && j < n; i++ {
		if nums[i]+nums[j] != n {
			if nums[i] != nums[i-1]+1 {
				return nums[i-1] + 1
			}
			if nums[j] != nums[j+1]-1 {
				return nums[j+1] - 1
			}
		}
	}
	return 0
}

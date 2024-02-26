package main

import "sort"

func minimumReplacement(nums []int) int64 {
	count := 0
	for !sort.IntsAreSorted(nums) {
		for idx, num := range nums {
			if num > nums[idx+1] {
				nums = replace(nums, idx, nums[idx-1], num-nums[idx-1])
				count++
			}
		}
	}
	return int64(count)
}

func replace(slice []int, index, newValue1, newValue2 int) []int {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = newValue1
	slice[index+1] = newValue2
	return slice
}

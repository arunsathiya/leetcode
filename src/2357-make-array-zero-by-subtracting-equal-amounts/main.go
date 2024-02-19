package main

import "sort"

func allZeroes(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}

func subtractByValue(slice []int, value int) []int {
	for i := range slice {
		if slice[i] > 0 {
			slice[i] -= value
		}
	}
	return slice
}

func minimumOperations(nums []int) int {
	count := 0
	for !allZeroes(nums) {
		sort.Ints(nums)
		var smallest int
		for _, v := range nums {
			if v > 0 {
				smallest = v
				break
			}
		}
		nums = subtractByValue(nums, smallest)
		count++
	}
	return count
}

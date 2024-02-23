package main

import "sort"

func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				output = append(output, []int{num, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
		}
	}
	return output
}

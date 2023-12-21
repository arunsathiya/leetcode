package main

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

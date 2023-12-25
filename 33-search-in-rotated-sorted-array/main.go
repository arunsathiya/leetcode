package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return left
		}
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return left
}

package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := (left + right) / 2
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

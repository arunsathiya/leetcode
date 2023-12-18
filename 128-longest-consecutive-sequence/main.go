package main

func sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func longestConsecutive(nums []int) int {
	nums = sort(nums)
	if len(nums) == 0 {
		return 0
	}
	maxLength, currentLength := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			currentLength++
		} else {
			currentLength = 1
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

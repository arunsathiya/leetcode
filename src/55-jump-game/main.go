package main

func canJump(nums []int) bool {
	lastIndex := len(nums) - 1
	for idx, num := range nums {
		if idx+num >= lastIndex {
			return true
		} else {
			continue
		}
	}
	return false
}

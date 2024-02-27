package main

func canJump(nums []int) bool {
	furthestReachable := 0
	lastIndex := len(nums) - 1
	for idx, num := range nums {
		if idx > furthestReachable {
			return false
		}
		furthestReachable = max(furthestReachable, idx+num)
		if furthestReachable >= lastIndex {
			return true
		}
	}
	return false
}

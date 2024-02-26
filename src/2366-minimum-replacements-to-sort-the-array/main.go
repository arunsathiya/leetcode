package main

func minimumReplacement(nums []int) int64 {
	count := 0
	for i := len(nums) - 1; i > 0; i-- {
		current := nums[i]
		prev := nums[i-1]
		if prev <= current {
			continue
		}
		operationsNeeded := (prev + current - 1) / current
		count += operationsNeeded - 1
		nums[i-1] = nums[i]
	}
	return int64(count)
}

package main

func containsDuplicate(nums []int) bool {
	hm := make(map[int]bool)
	for _, num := range nums {
		if hm[num] {
			return true
		}
		hm[num] = true
	}
	return false
}

package main

func containsDuplicate(nums []int) bool {
	hm := make(map[int]struct{})
	for _, num := range nums {
		if _, found := hm[num]; foun {
			return true
		}
		hm[num] = struct{}{}
	}
	return false
}

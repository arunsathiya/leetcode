package main

func intersection(nums1 []int, nums2 []int) []int {
	managed := []int{}
	hm := make(map[int]bool)
	for _, num1 := range nums1 {
		hm[num1] = true
	}
	for _, num2 := range nums2 {
		if hm[num2] {
			managed = append(managed, num2)
			hm[num2] = false
		}
	}
	return managed
}

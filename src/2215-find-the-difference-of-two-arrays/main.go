package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1hm := make(map[int]bool)
	nums2hm := make(map[int]bool)
	var nums1return []int
	var nums2return []int
	var final [][]int
	for _, num1 := range nums1 {
		nums1hm[num1] = true
	}
	for _, num2 := range nums2 {
		nums2hm[num2] = true
	}
	for _, num1 := range nums1 {
		if _, found := nums2hm[num1]; !found {
			if !existsInSlice(nums1return, num1) {
				nums1return = append(nums1return, num1)
			}
		}
	}
	for _, num2 := range nums2 {
		if _, found := nums1hm[num2]; !found {
			if !existsInSlice(nums2return, num2) {
				nums2return = append(nums2return, num2)
			}
		}
	}
	final = append(final, nums1return)
	final = append(final, nums2return)
	return final
}

func existsInSlice(nums []int, num int) bool {
	for _, number := range nums {
		if number == num {
			return true
		}
	}
	return false
}

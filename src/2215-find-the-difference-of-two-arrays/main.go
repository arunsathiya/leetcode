package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	return [][]int{setDifference(nums1, nums2), setDifference(nums2, nums1)}
}

func setDifference(nums1, nums2 []int) []int {
	indices := [2001]bool{}
	answer := []int{}
	for _, num2 := range nums2 {
		indices[num2+1000] = true
	}
	for _, num1 := range nums1 {
		if !indices[num1+1000] {
			answer = append(answer, num1)
			indices[num1+1000] = true
		}
	}
	return answer
}

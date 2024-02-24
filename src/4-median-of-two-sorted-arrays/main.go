package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	managed := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			managed = append(managed, nums1[i])
			i++
		} else {
			managed = append(managed, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		managed = append(managed, nums1[i])
		i++
	}
	for j < len(nums2) {
		managed = append(managed, nums2[j])
		j++
	}
	mid := len(managed) / 2
	if len(managed)%2 != 0 {
		return float64(managed[mid])
	}
	return (float64(managed[mid]) + float64(managed[mid-1])) / 2.0
}

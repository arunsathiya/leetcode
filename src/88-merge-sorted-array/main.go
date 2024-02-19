package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, mi := m-1, n-1, m+n-1
	for i2 >= 0 {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[mi] = nums1[i1]
			i1--
		} else {
			nums1[mi] = nums2[i2]
			i2--
		}
		mi--
	}
}

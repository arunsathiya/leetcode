package main

import "sort"

type Intslice []int

func (a Intslice) Len() int           { return len(a) }
func (a Intslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intslice) Less(i, j int) bool { return a[i] < a[j] }

func longestConsecutive(nums []int) int {
	sort.Sort(Intslice(nums))
	if len(nums) == 0 {
		return 0
	}
	maxLength, currentLength := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			currentLength++
		} else {
			if nums[i] != nums[i-1] {
				currentLength = 1
			}
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

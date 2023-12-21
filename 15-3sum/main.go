package main

import "sort"

type Intslice []int

func (a Intslice) Len() int           { return len(a) }
func (a Intslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intslice) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	sort.Sort(Intslice(nums))
	output := make([][]int, 0)
	for i := range nums {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				output = append(output, []int{nums[i], nums[left], nums[right]})
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return output
}

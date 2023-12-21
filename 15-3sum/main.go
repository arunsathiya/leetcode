package main

import "sort"

type Intslice []int

func (a Intslice) Len() int           { return len(a) }
func (a Intslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intslice) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	sort.Sort(Intslice(nums))
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				output = append(output, []int{num, nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return output
}

package main

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, num := range nums {
		if j, okay := hm[target-num]; okay {
			return []int{j, i}
		}
		hm[num] = i
	}
	return nil
}

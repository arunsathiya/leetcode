package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, num := range nums {
		if j, okay := hashmap[target-num]; okay {
			return []int{j, i}
		}
		hashmap[num] = i
	}
	return nil
}

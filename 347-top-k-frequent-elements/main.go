package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int, 0)
	results := make([]int, 0)
	for _, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = 1
		} else {
			hashmap[num] += 1
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return hashmap[nums[i]] > hashmap[nums[j]]
	})
	for _, num := range nums {
		if len(results) == 0 || results[len(results)-1] != num {
			results = append(results, num)
		}
		if len(results) == k {
			break
		}
	}
	return results
}

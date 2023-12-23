package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int, 0)
	occurances := make([]int, 0)
	results := make([]int, 0)
	for _, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = 1
		} else {
			hashmap[num] += 1
		}
	}
	for _, v := range hashmap {
		occurances = append(occurances, v)
	}
	sort.Slice(occurances, func(i, j int) bool {
		return occurances[i] > occurances[j]
	})
	for k, v := range hashmap {
		if v == occurances[0] || v == occurances[1] {
			results = append(results, k)
		}
	}
	return results
}

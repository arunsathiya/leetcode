package main

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int, 0)
	for _, num := range nums {
		if _, ok := hashmap[num]; !ok {
			hashmap[num] = 1
		} else {
			hashmap[num] += 1
		}
	}
	return nil
}

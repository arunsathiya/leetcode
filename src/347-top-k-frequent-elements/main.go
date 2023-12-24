package main

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
	invertedMap := make(map[int][]int, 0)
	maxFreq := 0
	for value, freq := range hashmap {
		invertedMap[freq] = append(invertedMap[freq], value)
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	for freq := maxFreq; freq > 0 && len(results) < k; freq-- {
		results = append(results, invertedMap[freq]...)
	}
	return results
}

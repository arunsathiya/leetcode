package main

func uniqueOccurrences(arr []int) bool {
	hm := make(map[int]int)
	for _, num := range arr {
		hm[num]++
	}
	values := make(map[int]int)
	for _, v := range hm {
		values[v]++
		if values[v] > 1 {
			return false
		}
	}
	return true
}

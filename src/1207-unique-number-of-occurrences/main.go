package main

func uniqueOccurrences(arr []int) bool {
	hm := make(map[int]int)
	for _, num := range arr {
		hm[num]++
	}
	values := make(map[int]bool)
	for _, v := range hm {
		if values[v] {
			return false
		}
		values[v] = true
	}
	return true
}

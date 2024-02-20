package main

import "strings"

func makeEqual(words []string) bool {
	hm := make(map[string]int)
	count := 0
	for _, word := range words {
		characters := strings.Split(word, "")
		count += len(characters)
		for _, char := range characters {
			hm[char]++
		}
	}
	if count%len(words) == 0 && areCharCountsEqual(hm) {
		return true
	}
	return false
}

func areCharCountsEqual(hm map[string]int) bool {
	var prevCount int
	for _, count := range hm {
		if prevCount == 0 {
			prevCount = count
		} else if prevCount != count {
			return false
		}
	}
	return true
}

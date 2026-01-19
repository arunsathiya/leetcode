package main

func makeEqual(words []string) bool {
	counter := make([]int, 26)
	for _, word := range words {
		for _, char := range word {
			counter[char-'a']++
		}
	}
	for _, count := range counter {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}

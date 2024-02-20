package main

func makeEqual(words []string) bool {
	hm := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			hm[char]++
		}
	}
	for _, count := range hm {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}

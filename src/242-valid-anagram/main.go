package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[rune]int)
	for _, char := range s {
		hm[char]++
	}
	for _, char := range t {
		if count, found := hm[char]; found && count > 0 {
			hm[char]--
		} else {
			return false
		}
	}
	return true
}

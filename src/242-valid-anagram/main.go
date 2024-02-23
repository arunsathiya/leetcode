package main

func isAnagram(s string, t string) bool {
	hm := make(map[rune]int)
	hmt := make(map[rune]int)
	for _, char := range s {
		hm[char]++
	}
	for _, char := range t {
		hmt[char]++
	}
	if !mapsEqual(hm, hmt) {
		return false
	}
	return true
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

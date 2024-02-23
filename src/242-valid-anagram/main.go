package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make([]int, 26)
	for _, char := range s {
		idx := int(char - 'a')
		chars[idx]++
	}
	for _, char := range t {
		idx := int(char - 'a')
		chars[idx]--
	}
	for _, char := range chars {
		if char != 0 {
			return false
		}
	}
	return true
}

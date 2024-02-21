package main

func partitionString(s string) int {
	optimal := make([]string, 0)
	for len(s) > 0 {
		instance := ""
		charMap := make(map[rune]bool)
		for _, char := range s {
			if charMap[char] {
				break
			}
			instance += string(char)
			charMap[char] = true
		}
		optimal = append(optimal, instance)
		s = s[len(instance):]
	}
	return len(optimal)
}

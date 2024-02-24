package main

func lengthOfLongestSubstring(s string) int {
	managed := ""
	chars := [26]int{}
	for _, char := range s {
		managed += string(char)
		idx := int(char - 'a')
		chars[idx]++
		if chars[idx] >= 2 {
			return len(managed) - 1
		}
	}
	return 0
}

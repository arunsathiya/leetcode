package main

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLength := 0
	charhm := make(map[rune]int)
	for end, char := range s {
		if trackedIdx, found := charhm[char]; found {
			start = trackedIdx + 1
		}
		charhm[char] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

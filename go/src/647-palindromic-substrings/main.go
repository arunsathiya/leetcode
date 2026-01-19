package main

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	var expandAroundCenter = func(left int, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if s[left] == s[right] {
				count++
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)
		expandAroundCenter(i, i+1)
	}
	return count
}

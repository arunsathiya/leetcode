package main

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	var expandAroundCenter = func(left int, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left) > (end - start) {
				start = left
				end = right
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)
		expandAroundCenter(i, i+1)
	}
	return s[start : end+1]
}

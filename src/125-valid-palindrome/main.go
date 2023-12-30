package main

import "unicode"

func isPalindrome(s string) bool {
	s = alphanumeric(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func alphanumeric(s string) string {
	var result []rune
	for _, r := range s {
		if ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || ('0' <= r && r <= '9') {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}

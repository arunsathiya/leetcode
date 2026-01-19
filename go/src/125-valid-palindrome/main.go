package main

func isPalindrome(s string) bool {
	managed := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if char >= 'A' && char <= 'Z' {
				char += 32
			}
			managed += string(char)
		}
	}
	left, right := 0, len(managed)-1
	for left < right {
		if managed[left] != managed[right] {
			return false
		}
		left++
		right--
	}
	return true
}

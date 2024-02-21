package main

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	palindromicSubstrings := make([]string, 0)
	for _, char := range s {
		palindromicSubstrings = append(palindromicSubstrings, string(char))
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			palindromicSubstrings = append(palindromicSubstrings, s[i:i+2])
		}
	}
	return len(palindromicSubstrings)
}

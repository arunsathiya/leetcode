package main

func mergeAlternately(word1 string, word2 string) string {
	var result string
	if len(word1) <= len(word2) {
		for i := 0; i < len(word1); i++ {
			result += string(word1[i])
			result += string(word2[i])
		}
		result += string(word2[len(word1):])
	} else {
		for i := 0; i < len(word2); i++ {
			result += string(word1[i])
			result += string(word2[i])
		}
		result += string(word1[len(word2):])
	}
	return result
}

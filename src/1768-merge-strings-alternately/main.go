package main

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result += string(word1[i])
		result += string(word2[i])
		i++
		j++
	}
	if i < len(word1) {
		result += string(word1[i:])
	}
	if j < len(word2) {
		result += string(word2[j:])
	}
	return result
}

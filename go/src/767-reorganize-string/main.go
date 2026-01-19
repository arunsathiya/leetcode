package main

func reorganizeString(s string) string {
	hm := make(map[rune]int)
	for _, char := range s {
		hm[char]++
	}
	var maxCount int
	var maxChar rune
	for k, v := range hm {
		if v > maxCount {
			maxCount = v
			maxChar = k
		}
	}
	// When no combination is possible
	if maxCount > (len(s)+1)/2 {
		return ""
	}
	// Fill even spaces
	result := make([]rune, len(s))
	index := 0
	for hm[maxChar] > 0 {
		result[index] = maxChar
		index += 2
		hm[maxChar]--
	}
	// Fill odd spaces
	for k, v := range hm {
		for v > 0 {
			if index >= len(s) {
				index = 1
			}
			result[index] = k
			index += 2
			v--
		}
	}
	return string(result)
}

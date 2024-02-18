package main

func reorganizeString(s string) string {
	hm := make(map[rune]int)
	for _, char := range s {
		hm[char]++
	}
	var maxCount int
	for _, v := range hm {
		if v > maxCount {
			maxCount = v
		}
	}
	if maxCount > (len(s)+1)/2 {
		return ""
	}
	return ""
}

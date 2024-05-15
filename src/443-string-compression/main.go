package main

import "strconv"

func compress(chars []byte) int {
	var result int
	hm := make(map[byte]int)
	for _, char := range chars {
		hm[char]++
	}
	for _, v := range hm {
		if v == 1 {
			result++
		}
		if v > 1 && v < 10 {
			result += 1
			result += v
		}
		if v > 10 {
			result += 1
			result += len(strconv.Itoa(v))
		}
	}
	return result
}

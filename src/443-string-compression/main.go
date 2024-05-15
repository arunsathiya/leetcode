package main

import "strconv"

func compress(chars []byte) int {
	original := len(chars)
	hm := make(map[byte]int)
	for _, char := range chars {
		hm[char]++
	}
	for char, count := range hm {
		chars = append(chars, char)
		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars = append(chars, byte(digit))
			}
		}
	}
	return len(chars) - original
}

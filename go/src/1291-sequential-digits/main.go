package main

import (
	"fmt"
	"strconv"
)

func sequentialDigits(low int, high int) []int {
	results := make([]int, 0)
	base := "123456789"
	lowl, highl := len(fmt.Sprintf("%d", low)), len(fmt.Sprintf("%d", high))
	for length := lowl; length <= highl; length++ {
		for start := 0; start <= 9-length; start++ {
			substring := base[start : start+length]
			num, _ := strconv.Atoi(substring)
			if num >= low && num <= high {
				results = append(results, num)
			}
		}
	}
	return results
}

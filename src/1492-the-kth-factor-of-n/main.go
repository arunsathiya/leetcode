package main

import "sort"

func kthFactor(n int, k int) int {
	factors := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if _, exists := factors[i]; !exists {
				factors[i] = true
				factors[n/i] = true
			} else {
				break
			}
		}
	}
	var s []int
	for k := range factors {
		s = append(s, k)
	}
	sort.Ints(s)
	if k <= len(s) {
		return s[k-1]
	} else {
		return -1
	}
}

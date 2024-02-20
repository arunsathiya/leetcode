package main

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
	sort(s)
	if k <= len(s) {
		return s[k-1]
	} else {
		return -1
	}
}

func sort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

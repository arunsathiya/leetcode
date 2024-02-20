package main

func kthFactor(n int, k int) int {
	factors := make([]int, 0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
		if len(factors) == k {
			return factors[k-1]
		}
	}
	factors = append(factors, n)
	if k <= len(factors) {
		return factors[k-1]
	} else {
		return -1
	}
}

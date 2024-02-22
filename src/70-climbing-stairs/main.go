package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ways := make([]int, n+1)
	for i := 3; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}

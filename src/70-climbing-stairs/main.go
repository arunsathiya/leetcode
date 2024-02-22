package main

func climbStairs(n int) int {
	memo := make(map[int]int, 0)
	return climbStairsMemo(n, memo)
}

func climbStairsMemo(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if _, exists := memo[n]; exists {
		return memo[n]
	}
	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	return memo[n]
}

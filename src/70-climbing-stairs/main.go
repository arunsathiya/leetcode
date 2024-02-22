package main

func climbStairs(n int) int {
	memo := make([]int, 46)
	return climbStairsMemo(n, memo)
}

func climbStairsMemo(n int, memo []int) int {
	if n <= 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)
	return memo[n]
}

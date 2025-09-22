package main

func reverse(x int) int {
	var remainder int
	var remaining int
	var result int
	const MaxInt32 = 2147483647
	const MinInt32 = -2147483648
	originalX := x
	if x < 0 {
		x = -x
	}
	remainder = x % 10
	remaining = x / 10
	result += (result * 10) + int(remainder)
	for remaining > 0 {
		remainder = remaining % 10
		remaining = remaining / 10
		if result > (MaxInt32-remainder)/10 {
			return 0
		}
		if result < (MinInt32-remainder)/10 {
			return 0
		}
		result = (result * 10) + remainder
	}
	if originalX < 0 {
		return -result
	}
	return result
}

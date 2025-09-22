package main

func reverse(x int) int {
	var remainder int
	var result int
	const MaxInt32 = 2147483647
	const MinInt32 = -2147483648
	originalX := x
	if x < 0 {
		x = -x
	}
	for x > 0 {
		remainder = x % 10
		x = x / 10
		if result > (MaxInt32-remainder)/10 {
			return 0
		}
		if originalX < 0 && result > (-MinInt32-remainder)/10 {
			return 0
		}
		result = (result * 10) + remainder
	}
	if originalX < 0 {
		return -result
	}
	return result
}

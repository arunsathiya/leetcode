package main

func reverse(x int) int {
	var remainder int
	var remaining int
	var result int
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
		result = (result * 10) + remainder
	}
	if originalX < 0 {
		return -result
	}
	return result
}

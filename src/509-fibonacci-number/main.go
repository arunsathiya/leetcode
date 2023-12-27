package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}

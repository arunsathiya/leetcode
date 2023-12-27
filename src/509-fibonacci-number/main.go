package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	sequence := make([]int, 0)
	sequence = append(sequence, 0)
	sequence = append(sequence, 1)
	for i := 2; i < n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}
	return sequence[n-1] + sequence[n-2]
}

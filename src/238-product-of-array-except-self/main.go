package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	output := make([]int, n)
	output[0] = 1

	for i := 1; i < n; i++ {
		output[i] = output[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		output[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return output
}

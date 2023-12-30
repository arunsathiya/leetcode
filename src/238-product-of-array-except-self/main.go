package main

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	output := make([]int, len(nums))
	output[0] = 1
	for i := 1; i < len(nums); i++ {
		output[0] *= nums[i]
	}
	output[len(nums)-1] = 1
	for i := 0; i < len(nums)-1; i++ {
		output[len(nums)-1] *= nums[i]
	}
	for i := 1; i < len(nums)-1; i++ {
		leftProduct, rightProduct := 1, 1
		for j := 0; j < i; j++ {
			leftProduct *= nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			rightProduct *= nums[j]
		}
		output[i] = leftProduct * rightProduct
	}
	return output
}

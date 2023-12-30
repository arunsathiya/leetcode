package main

func productExceptSelf(nums []int) []int {
	if status, indices := containsZero(nums); status {
		for index, num := range nums {
			for zeroIndex, _ := range indices {
				if index == zeroIndex {
					nums[index] = 
				}
			}
		}
	}
}

func containsZero(nums []int) (bool, []int) {
	indices := make([]int,0)
	for i, num := range nums {
		if num == 0 {
			indices = append(indices, i)
			return true, indices
		}
	}
	return false, nil
}

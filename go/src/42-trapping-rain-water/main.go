package main

func trap(height []int) int {
	totalTrap := 0
	highestIndex := 0
	for idx, h := range height {
		if h > height[highestIndex] {
			highestIndex = idx
		}
	}
	leftMax := -1
	for i := 0; i < highestIndex; i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		totalTrap += leftMax - height[i]
	}
	rightMax := -1
	for i := len(height) - 1; i > highestIndex; i-- {
		if height[i] > rightMax {
			rightMax = height[i]
		}
		totalTrap += rightMax - height[i]
	}
	return totalTrap
}

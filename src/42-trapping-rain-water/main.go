package main

func trap(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	totalTrap := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalTrap += leftMax - height[left]
			}
			left++
		}
		if height[right] <= height[left] {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalTrap += rightMax - height[right]
			}
			right--
		}
	}
	return totalTrap
}

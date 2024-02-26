package main

func trap(height []int) int {
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	leftTrap, rightTrap := 0, 0
	totalTrap := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] < leftMax {
				leftTrap = leftMax - height[left]
			}
			totalTrap += leftTrap
			if height[left] > leftMax {
				leftMax = height[left]
			}
			left++
		}
		if height[right] <= height[left] {
			if height[right] < rightMax {
				rightTrap = rightMax - height[right]
			}
			totalTrap += rightTrap
			if height[right] > rightMax {
				rightMax = height[right]
			}
			right--
		}
	}
	return totalTrap
}

package main

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {
	var area int
	var maxArea int
	var diagonal float64
	var maxDiagonal float64
	for i := range dimensions {
		length := dimensions[i][0]
		width := dimensions[i][1]
		diagonal = math.Sqrt(float64(length)*float64(length) + float64(width)*float64(width))
		area = length * width
		if diagonal > maxDiagonal {
			maxDiagonal = diagonal
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

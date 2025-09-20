package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

func areaOfMaxDiagonal(dimensions string) int {
	var area int
	var maxArea int
	var diagonal float64
	var maxDiagonal float64
	var parsedDimensions [][]int
	err := json.Unmarshal([]byte(dimensions), &parsedDimensions)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range parsedDimensions {
		length := parsedDimensions[i][0]
		width := parsedDimensions[i][1]
		diagonal = math.Sqrt(float64(length)*float64(length) + float64(width)*float64(width))
		area = length * width
		fmt.Printf("Rectangle %d: [%d,%d], diagonal=%.2f, area=%d\n", i, length, width, diagonal, area)
		if diagonal > maxDiagonal {
			fmt.Printf(" -> NEW MAX! Previous: %2.f, New: %.2f\n", maxDiagonal, diagonal)
			maxDiagonal = diagonal
			maxArea = area
		}
		if diagonal == maxDiagonal {
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

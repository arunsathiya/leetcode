package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, (rows*cols)-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/cols][mid%cols]
		if target == midValue {
			return true
		}
		if target > midValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

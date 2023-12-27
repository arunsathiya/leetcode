package main

func searchMatrix(matrix [][]int, target int) bool {
	firstmatrix := matrix[0]
	left, right := 0, len(firstmatrix)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == firstmatrix[mid] {
			return true
		}
		if firstmatrix[mid] >= firstmatrix[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

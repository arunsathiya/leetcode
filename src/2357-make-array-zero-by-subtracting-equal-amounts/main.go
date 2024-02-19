package main

func minimumOperations(nums []int) int {
	tmp := make(map[int]int)
	for _, num := range nums {
		if num != 0 {
			tmp[num]++
		}
	}
	return len(tmp)
}

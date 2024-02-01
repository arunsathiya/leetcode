package main

func platesBetweenCandles(s string, queries [][]int) []int {
	within := []int{}
	for _, query := range queries {
		ss := s[query[0]:query[1]]
		count, inContainer := 0, false
		for _, character := range ss {
			switch character {
			case '|':
				inContainer = !inContainer
			case '*':
				if inContainer {
					count++
				}
			}
		}
		within = append(within, count)
	}
	return within
}

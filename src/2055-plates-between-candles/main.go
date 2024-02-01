package main

import "sort"

func platesBetweenCandles(s string, queries [][]int) []int {
	candles := make([]int, 0, len(s))
	for i, character := range s {
		if character == '|' {
			candles = append(candles, i)
		}
	}
	answer := make([]int, 0, len(queries))
	for _, query := range queries {
		l := query[0]
		r := query[1]
		left := sort.Search(len(candles), func(i int) bool {
			return candles[i] >= l
		})
		right := sort.Search(len(candles), func(i int) bool {
			return candles[i] > r
		}) - 1
		if right <= left {
			answer = append(answer, 0)
			continue
		}
		c := (candles[right] - candles[left]) - (right - left)
		answer = append(answer, c)
	}
	return answer
}

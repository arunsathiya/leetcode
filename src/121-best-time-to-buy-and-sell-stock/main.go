package main

func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0
	for right < len(prices) {
		currentProfit := prices[right] - prices[left]
		if prices[right] > prices[left] {
			maxProfit = Max(currentProfit, maxProfit)
		} else {
			left = right
		}
		right += 1
	}
	return maxProfit
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

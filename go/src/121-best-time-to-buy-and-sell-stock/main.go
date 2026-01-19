package main

func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0
	for right < len(prices) {
		currentProfit := prices[right] - prices[left]
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
		if prices[right] < prices[left] {
			left = right
		}
		right++
	}
	return maxProfit
}

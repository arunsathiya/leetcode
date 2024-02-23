package main

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	currentProfit := 0
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			currentProfit = prices[i] - buyPrice
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}
	return maxProfit
}

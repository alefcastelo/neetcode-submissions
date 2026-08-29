func maxProfit(prices []int) int {
	buye, sell, profit := 0, 1, 0

	for sell < len(prices) {
		if prices[buye] >= prices[sell] {
			buye = sell
			sell++ 
			continue
		}

		currentProfit := prices[sell] - prices[buye]

		if currentProfit > profit {
			profit = currentProfit
		}

		sell++
	}

	return profit
}

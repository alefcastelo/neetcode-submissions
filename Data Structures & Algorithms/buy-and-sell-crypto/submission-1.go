func maxProfit(prices []int) int {
	profit, buy, sell := 0, 0, 1

	for sell < len(prices) {
		if prices[sell] - prices[buy] > profit {
			profit = prices[sell] - prices[buy]
		}

		if prices[sell] < prices[buy] {
			buy++
		} else {
			sell++
		}
	}

	return profit
}

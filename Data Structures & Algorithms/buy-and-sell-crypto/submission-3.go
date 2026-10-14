func maxProfit(prices []int) int {
	profit := 0

	buy, sell := 0, 1

	for sell < len(prices) {

		if prices[buy] > prices[sell] {
			buy = sell
			sell++
			continue
		}

		diff := prices[sell] - prices[buy]

		if diff > profit {
			profit = diff
		}

		sell++
	}


	return profit
}

// buy    7
// sell   1
// profit 

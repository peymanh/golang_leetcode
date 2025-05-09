package buyandsell

func MaxProfit(prices []int) int {
	profit, buy := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i]-buy > profit {
			profit = prices[i] - buy
		}
	}
	return profit

}

package buyandsell2

func Accumulate(elements ...int) int {
	sum := 0
	for _, elem := range elements {
		sum += elem
	}
	return sum
}

func maxProfit(prices []int) int {

	profit := make([]int, 0)

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {
			profit = append(profit, (prices[i+1] - prices[i]))
		}
	}

	return Accumulate(profit...)
}

func maxProfit(prices []int) int {
	maxProfit := 0

	for i:=0; i<len(prices)-1; i++ {
		profit := 0
		for j:=i+1; j<len(prices); j++ {
			tmp := prices[j] - prices[i]

			if tmp > profit {
				profit = tmp
			}
		}

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func maxProfit(prices []int) int {
	maxProfit := 0

	left, right := 0, 1

	for right < len(prices) {
		if prices[right] <= prices[left] {
			left=right
		} else {
			diff := prices[right] - prices[left]
			if diff > maxProfit {
				maxProfit = diff
			}
		}
		right++
	}

	return maxProfit
}

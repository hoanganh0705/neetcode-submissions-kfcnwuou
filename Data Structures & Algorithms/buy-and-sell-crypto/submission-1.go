func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	offset := 0
	for i:=1; i< len(prices); i++ {
		if minPrice > prices[i] {
			minPrice =  prices[i]
		}

		if (prices[i] - minPrice) > offset {
			offset = prices[i] - minPrice
		}
	}

	return offset
}

package greedy

// 买卖股票的最佳时机Ⅱ: 贪心,只要这一次比上一次大,就"买上",本质上在一个向上的坡段里,在最低处买,在最高处卖
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			res += prices[i] - prices[i - 1]
		}
	}

	return res
}


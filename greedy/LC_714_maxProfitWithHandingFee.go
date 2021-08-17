package greedy

// 买卖股票含手续费: 类似Ⅱ,核心思想:当我们卖出一只股票时,我们便拥有了将其再买入,且卖出时不给手续费的权利
func maxProfit4(prices []int, fee int) int {
	res := 0

	m := prices[0]
	for i := 1; i < len(prices); i++ {
		m = min(m, prices[i])
		if prices[i] - m > fee {
			res += prices[i] - m - fee
			m = prices[i] - fee
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


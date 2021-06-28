package gaming_theory

//本质：区间DP
//博弈论：数学相关可直接求解
//可进行状态压缩，未来需补充
func stoneGame(piles []int) bool {
	m := len(piles)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		dp[i][i] = piles[i]
	}

	for i := m - 1; i >= 0; i-- {
		for  j := i + 1; j < m; j++ {
			dp[i][j] = max(piles[i] - dp[i + 1][j], piles[j] - dp[i][j - 1])
		}
	}

	if dp[0][m - 1] > 0 {
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

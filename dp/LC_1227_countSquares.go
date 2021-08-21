package dp

// 统计全为一的正方形子矩阵
func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i - 1][j - 1] == 1 {
				dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
				res += dp[i][j]
			}
		}
	}

	return res
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

package dp

// 最大正方形 -> 注:求最大面积,得到的最大边长还需要再平方
// dp[i][j]含义: (i, j)处为右下角的正方形的最小边长 -> 受其左,上,左上三者的影响,取最小值即可
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i - 1][j - 1] == '1' {
				dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}

	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

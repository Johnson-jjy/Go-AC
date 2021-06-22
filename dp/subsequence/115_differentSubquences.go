package subsequence

// 未来补充：滚动数组压缩
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1]
			} else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}

	return dp[m][n]
}

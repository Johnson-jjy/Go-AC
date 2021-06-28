package subsequence

// 注1：此处dp数组的定义，表示长度
// 注2：对于else--本质是对t的“编辑删除”
// 未完待续：后续挑战！
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)

	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = dp[i][j - 1]
			}
		}
	}

	return dp[m][n] == len(s)
}

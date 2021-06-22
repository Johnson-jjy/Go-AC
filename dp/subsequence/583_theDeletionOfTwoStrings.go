package subsequence

// 动归1：利用LCS，求出两字符串的最长公共子序列后，再分别删除即可
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
			}
		}
	}

	return m + n - dp[m][n] - dp[m][n]
}

// 动归2：直接定义dp数组表示删除所需要的次数
// 注：初始化一定要留意
func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m + 1)
	dp[0] = make([]int, n + 1)
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n + 1)
		dp[i][0] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				cur := min(dp[i - 1][j] + 1, dp[i][j - 1] + 1)
				dp[i][j] = min(cur, dp[i - 1][j - 1] + 2)
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
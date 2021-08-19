package subsequence

// 编辑距离: 经典dp[i][j], 以i结尾的字串1和以j结尾的字串2, 最少需要的操作数
// 注: 1.初始化的时候对各种0的情况初始化勿忘; 2.对于该处字符不等的情况,本质是 1)修改 2)删i 3)删j
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m + 1)
	dp[0] = make([]int, n + 1)
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		dp[i] = make([]int,n+1)
		dp[i][0] = i
		for j := 1; j <= n; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(min(dp[i][j - 1] + 1, dp[i - 1][j] + 1), dp[i - 1][j - 1] + 1)
			}
		}
	}

	return dp[m][n]
}

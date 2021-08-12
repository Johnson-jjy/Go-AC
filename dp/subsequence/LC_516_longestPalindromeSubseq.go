package subsequence

// 注: 子序列的题目一般喜欢多开一个位置 -> 遍历顺序与传统的不同,需要在有了递推公式后进一步判断
func longestPalindromeSubseq(s string) int {
	m := len(s)
	dp := make([][]int, m + 1)
	//1.初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, m + 1)
		dp[i][i] = 1
	}
	//2.遍历顺序
	for i := m; i >= 1; i-- {
		for j := i + 1; j <= m; j++ {
			// 3.递推关系
			if s[i - 1] == s[j - 1] {
				dp[i][j] = dp[i + 1][j - 1] + 2
			} else {
				dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}

	return dp[1][m]
}

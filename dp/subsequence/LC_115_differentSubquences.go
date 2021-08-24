package subsequence

// 不同的子序列 -> 1.理解dp递推方程的缘由与含义; 2.勿忘初始化
// 未来补充：滚动数组压缩
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
		dp[i][0] = 1 // 空串一定能够被匹配,本质上也是一种初始化的原有
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == t[j - 1] { // 相等的时候,可选择使用i或不使用i, 由两个结果来
				dp[i][j] = dp[i - 1][j] + dp[i - 1][j - 1]
			} else { // 不等的时候只能不使用i, 因为j必须全匹配,故不能j-1
				dp[i][j] = dp[i - 1][j]
			}
		}
	}

	return dp[m][n]
}

package subsequence

// 待补充一：中心扩散法
// 待补充二：manacher
// 可优化：压缩数组
func countSubstrings(s string) int {
	m := len(s)
	res := m
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, m + 1)
		dp[i][i] = true
	}

	for i := m; i >= 1; i-- {
		for j:= i + 1; j <= m; j++ {
			if s[i - 1] == s[j - 1] {
				if j - i <= 1 {
					dp[i][j] = true
					res++
				} else {
					dp[i][j] = dp[i + 1][j - 1]
					if dp[i][j] {
						res++
					}
				}
			}
		}
	}

	return res
}

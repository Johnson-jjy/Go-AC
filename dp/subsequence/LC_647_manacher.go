package subsequence

// 回文子串
// 待补充一：中心扩散法
// 待补充二：manacher
// 可优化：压缩数组
// 思路: 基础dp -> 回文字串类题目, 主要使用bool数组,对以i开头,j结尾的字符串进行判断,其是否为回文串
func countSubstrings(s string) int {
	m := len(s)
	dp := make([][]bool, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, m + 1)
		dp[i][i] = true
	}

	res := 0
	for i := m - 1; i >= 1; i-- {
		for j := i + 1; j <= m; j++ {
			if s[i - 1] == s[j - 1] {
				if i + 1 == j { // 相邻
					dp[i][j] = true
					res++
				} else {
					if dp[i + 1][j - 1] {
						dp[i][j] = true
						res++
					}
				}
			}
		}
	}

	return res + m
}
package dp

// 解码方法
// 有点类似于正则匹配类dp,关键是理清楚当前字符为某值时,对应的递推转换情况
// 注: 当前值时0时的讨论
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1]
		if s[i - 1] == '0' {
			if s[i - 2] != '1' && s[i - 2] != '2' {
				return 0
			} else {
				dp[i] = dp[i - 2]
			}
		} else if s[i - 2] == '1' || (s[i - 2] == '2' && s[i - 1] <= '6' && s[i - 1] >= '0') {
			dp[i] += dp[i - 2]
		}
	}
	return dp[n]
}

package dp

// 通配符匹配: 类似正则表达式匹配
// 难点1: 对于*的处理 -> 对与*而言,dp转移有两种,一种是不必配,一种是匹配一个或多个,对于多个的情况,直接i-1即可
// 难点2: 对于p开头多个*的处理 -> 需要处理的是多个*都不匹配的清类,即dp[0][i] = true -> 匹配多个的情况会在之后的dp转移中求出
// 注: 开头的if判断其实也是可以不需要的

// 另注: 贪心待补充 + 甚至是AC自动机
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	// if m == 0 && n == 0 {
	//     return true
	// }
	// if m > 0 && n == 0 {
	//     return false
	// }
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true
	for i := 0; i < n && p[i] == '*'; i++ {
		dp[0][i + 1] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == p[j - 1] || p[j - 1] == '?' {
				dp[i][j] = dp[i - 1][j - 1]
			} else if p[j - 1] == '*' {
				dp[i][j] = dp[i - 1][j] || dp[i][j - 1]
			}
		}
	}

	return dp[m][n]
}

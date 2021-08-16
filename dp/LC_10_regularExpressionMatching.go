package dp

// 正则表达式匹配: dp匹配类母题
// 难点1: 对于 ?*类的匹配 -> 匹配0个, 一个, 多个 -> 分别对应不同的表达式
// 难点2: 对于p开头的?*都不匹配的特殊初始化 -> s需要从0(即空串)开始枚举 ->于是相应的需要对i进行一些特殊判断 -> 整个if的写法顺序也需要有所侧重
// 针对难点2的测试用例: "aaabaaaababcbccbaab" -- "c*c*.*c*a*..*c*"
func isMatch10(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j - 1] == '*' { // 注意别漏'.'的情况, 同时||相关不能忘记括号
				if i > 0 && (s[i - 1] == p[j - 2] || p[j - 2] == '.') {
					dp[i][j] = dp[i - 1][j] || dp[i][j - 1] || dp[i][j - 2]
				} else {
					dp[i][j] = dp[i][j - 2]
				}
			} else if i == 0 { // 非p开头不匹配的情况,一定要注意防止字符越界危险
				continue
			} else if s[i - 1] == p[j - 1] || p[j - 1] == '.' {
				dp[i][j] = dp[i - 1][j - 1]
			}
		}
	}

	return dp[m][n]
}

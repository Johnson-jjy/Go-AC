package dp

func longestValidParentheses(s string) int {
	m := len(s)
	res := 0
	dp := make([]int, m + 1) // 以 i-1 为结尾的字符串的最长有效匹配

	for i := 2; i <= m; i++ {
		// 若未左括号则以该下标为结束的dp数组必然为0
		if s[i - 1] == ')' {
			// 注意，i从2开始那么必然有 i > 2
			if s[i - 2] == '(' {
				dp[i] = dp[i - 1] + 2 +dp[i - 2]
			} else {
				// 这列的坐标计算需要很细致
				t := dp[i - 1]
				if i - 1 - t > 0 && s[i - 2 - t] == '(' {
					if i - 1 - t > 2 {
						dp[i] = dp[i - 1] + 2 + dp[i - t - 2]
					} else {
						dp[i] = dp[i - 1] + 2
					}
				}
			}
			if res < dp[i] {
				res = dp[i]
			}
		}
	}

	return res
}

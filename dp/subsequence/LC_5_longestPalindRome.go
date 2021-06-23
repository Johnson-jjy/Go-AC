package subsequence

func longestPalindrome(s string) string {
	res := 1
	m := len(s)
	dp := make([][]bool, m + 1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, m + 1)
		dp[i][i] = true
	}

	start := 0
	end := 0

	for i := m; i >= 1; i-- {
		for j := i + 1; j <= m; j++ {
			if s[i - 1] == s[j - 1] {
				// 相邻的情况
				if j - i == 1 {
					dp[i][j] = true
					if res == 1 {
						res = 2
						start = i - 1
						end = j - 1
					}
				} else {
					dp[i][j] = dp[i + 1][j - 1]
					if dp[i][j] {
						// 如果是更长的子串，则更新始末位置
						if j - i + 1 > res {
							res = j - i + 1
							start = i - 1
							end = j - 1
						}
					}
				}
			}
		}
	}

	return string(s[start:end + 1])
}

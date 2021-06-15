package subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1) + 1)
	for i := 0; i < len(text1) + 1; i++ {
		dp[i] = make([]int, len(text2) + 1)
	}
	res := 0
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i + 1][j + 1] = max(dp[i][j] + 1, dp[i + 1][j + 1])
			} else {
				dp[i + 1][j + 1] = max(dp[i + 1][j], dp[i][j + 1])
			}
			res = max(dp[i + 1][j + 1], res)
		}
	}

	return res
}


package offer

// 动归:dp[i]表示绳长为i时可得的最大值
// 注意对于dp[i]和i大小的的讨论
// 数学方法待补充
func cuttingRope(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			cur := max(dp[j], j) * max(dp[i - j], i - j)
			if dp[i] < cur {
				dp[i] = cur
			}
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

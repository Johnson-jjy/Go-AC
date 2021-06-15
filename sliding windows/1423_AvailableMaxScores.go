package sliding_windows

// 此处使用的滑动数组和前缀数组；也可以不用前缀数组，直接使用滑动数组，本质相同，少开一个slice可更快
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	dp := make([]int, n + 1)
	sum := 0
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
		dp[i + 1] = sum
	}
	if n == k {
		return sum
	}

	window := dp[n - k]
	for i := n - k + 1; i <= n; i++ {
		window = min(window, dp[i] - dp[i + k - n])
	}

	return sum - window
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

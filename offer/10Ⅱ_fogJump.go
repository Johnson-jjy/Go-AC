package offer

func numWays(n int) int {
	const mod = 1000000007
	dp := make([]int, 2)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i % 2 == 0 {
			dp[i % 2] = (dp[i % 2] + dp[i % 2 + 1]) % mod
		} else {
			dp[i % 2] = (dp[i % 2] + dp[i % 2 - 1]) % mod
		}
	}
	if n % 2 == 0 {
		return dp[0]
	}
	return dp[1]
}

package offer

// 注意 i%2 为0或1时的细微不同处理
func fib(n int) int {
	const mod = 1000000007
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i % 2 == 0 {
			dp[i % 2] = (dp[i % 2] + dp[i % 2 + 1])%mod
		} else {
			dp[i % 2] = (dp[i % 2] + dp[i % 2 - 1])%mod
		}
	}
	if n % 2 == 0 {
		return dp[0]
	}
	return dp[1]
}

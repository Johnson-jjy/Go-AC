package array

// 获取生成数组中的最大值
// 注: 当无法找到明显规律,而题目数据又给的不大时,便可直接模拟 -> 模拟的时候注意是i不是n!
// 模拟时,也可有一些细节,缩小搜索范围 -> 如res只可能在index为奇数时出现, 只需要搜索一般的空间等
func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}
	res := 1
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i % 2 == 0 {
			dp[i] = dp[i / 2]
		} else {
			dp[i] = dp[i / 2] + dp[i / 2 + 1]
			if res < dp[i] {
				res = dp[i]
			}
		}
	}
	//fmt.Println(dp)
	return res
}

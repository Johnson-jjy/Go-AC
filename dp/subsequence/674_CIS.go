package subsequence

// 用了dp，多开了切片，直接贪心更快
func findLengthOfLCIS(nums []int) int {
	m := len(nums)
	dp := make([]int, m)
	res := 0
	for i := 1; i < m; i++ {
		if nums[i] > nums[i - 1] {
			dp[i] = dp[i - 1] + 1
			res = max(res, dp[i])
		}
	}

	return res + 1
}


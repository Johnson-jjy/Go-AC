package subsequence

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums) + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i + 1] = max(dp[j + 1] + 1, dp[i + 1])
			}
			res = max(res, dp[i + 1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

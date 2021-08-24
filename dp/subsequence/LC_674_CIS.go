package subsequence

// 最长连续递增序列
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

// 贪心 -> 每次记录一下cur子序列的curLen, 若更长, 则更新
func findLengthOfLCIS2(nums []int) int {
	m := len(nums)
	cur := 1

	res := 1
	for i := 1; i < m; i++ {
		if nums[i] > nums[i - 1] {
			cur++
			if cur > res {
				res = cur
			}
		} else {
			cur = 1
		}
	}

	return res
}

package subsequence

// 不相交的线
// 认清本质: 就是一道最长连续公共子序列
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1) + 1, len(nums2) + 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i + 1][j + 1] = dp[i][j] + 1
			} else {
				dp[i + 1][j + 1] = max(dp[i + 1][j], dp[i][j + 1])
			}
		}
	}
	return dp[m][n]
}


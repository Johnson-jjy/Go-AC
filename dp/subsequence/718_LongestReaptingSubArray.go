package subsequence

// 法2：滑动窗口，待补充
// 法3：二分查找 + 双指针 待补充
func findLength(nums1 []int, nums2 []int) int {
	len1 := len(nums1) + 1
	len2 := len(nums2) + 1
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}

	res := 0
	for i := 0; i < len1 - 1; i++ {
		for j := 0; j < len2 - 1; j++ {
			if nums1[i] == nums2[j] {
				dp[i + 1][j + 1] = dp[i][j] + 1
				res = max(res, dp[i + 1][j + 1])
			}
		}
	}
	return res
}

package subsequence

// 最长重复子数组
// 法1:dp-> 注意区别子数组和子序列,明确dp[i][j]的含义是,以当前两个字符结尾的公共子数组长度
// 因为指标是以改字符结尾的公共子数组的长度, res是每一次都需要重新计数的
// 法2:滑动窗口，待补充
// 法3:二分查找 + 双指针 待补充
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

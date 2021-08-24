package subsequence

// 最大子序和
// 注释的为dp数组最直观地写法，未注释的为压缩数组后的写法
// 线段树相关的写法未来补充
func maxSubArray(nums []int) int {
	m := len(nums)
	//dp := make([]int, m)
	//dp[0] = nums[0]
	//res := dp[0]
	res := nums[0]
	pre := nums[0]
	for i := 1; i < m; i++ {
		//dp[i] = max(dp[i - 1] + nums[i], nums[i])
		pre = max(nums[i], pre + nums[i])
		res = max(res, pre)
	}

	return res
}

// 0824
func maxSubArray2(nums []int) int {
	res := nums[0]
	dp := make([]int, len(nums) + 1)

	for i := 1; i <= len(nums); i++ {
		dp[i] = nums[i - 1] // dp[i]可以为当前nums的值
		if dp[i - 1] > 0 { // 若前一个值>0,才有相加的意义
			dp[i] += dp[i - 1]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
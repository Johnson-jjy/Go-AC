package array

// 一维数组的动态和
// 解一: 前缀和思想
func runningSum1(nums []int) []int {
	preSum := make([]int, len(nums) + 1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i - 1] + nums[i - 1]
	}

	return preSum[1:]
}

// 解二: 直接原数组修改即可
func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
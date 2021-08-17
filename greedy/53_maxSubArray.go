package greedy

import "math"

// 最大子序和: 注意res定义为最小,每次先+sum, 判一次res(防止数组中每一个数都是负数,去最大的那个负数)
// 另: 动态规划; 分治
func maxSubArray(nums []int) int {
	sum := 0
	res := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return res
}

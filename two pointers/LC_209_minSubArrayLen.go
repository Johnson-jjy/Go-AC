package two_pointers

import "math"

// 长度最小的子数组: 对于只考虑一个数组的情况, 可以使用双指针
// 注意对res计数的位置 -> 在每一次可移动left前计数
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	left := 0
	right := 0
	sum := 0
	for right < len(nums) {
		if nums[right] >= target {
			return 1
		}
		for right < len(nums) && sum < target {
			sum += nums[right]
			right++
		}
		for left < right && sum >= target {
			if right - left < res {
				res = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}

	return res
}

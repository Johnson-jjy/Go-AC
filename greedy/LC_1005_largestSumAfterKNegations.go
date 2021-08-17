package greedy

import "sort"

// K次取反后最大化的数组和
// 1. 要先排序,想办法把大的负数翻转过来  2.没必要取纠结k和负数的关系,直接先翻转大的负数即可
// 3. 对于index定位,初始值应该考量的是不走else的情况
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	index := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		} else {
			index = i
			break
		}
	}

	if k > 0 {
		if k % 2 == 1 {
			m := index
			if index > 0 && nums[index - 1] < nums[index] {
				m = index - 1
			}
			nums[m] = -nums[m]
		}
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}

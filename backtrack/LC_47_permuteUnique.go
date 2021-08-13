package backtrack

import "sort"

// 全排列2: 1.不能含重复->used数组判重的同时,对同层的相同字符进行跳过处理(continue)
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	cur := make([]int, 0)

	var backtrack func(cur []int)
	backtrack = func(cur []int) {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
				continue
			}
			if !used[i] {
				cur = append(cur, nums[i])
				used[i] = true
				backtrack(cur)
				cur = cur[:len(cur) - 1]
				used[i] = false
			}
		}
	}

	backtrack(cur)

	return res
}

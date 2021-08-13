package backtrack

import "sort"

// 组合2: 1.依然是要求最后的结果故不按背包处理; 2.有重复数字 -> 必涉及used判重
// 3.used[i-1] == false -> 同层判重更快; 4.注意,判重的情况是continue而不是return
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)

	if target < candidates[0] {
		return res
	}

	used := make([]bool, len(candidates))

	var backtrack func(cur []int, sum int, index int)
	backtrack = func(cur []int, sum int, index int) {
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if sum + candidates[i] > target {
				return
			}
			if i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1] {
				continue // 注意此处是continue, 而不是return
			}
			cur = append(cur, candidates[i])
			sum += candidates[i]
			used[i] = true
			backtrack(cur, sum, i + 1)
			cur = cur[:len(cur) - 1]
			sum -= candidates[i]
			used[i] = false
		}
	}
	cur := make([]int, 0)
	backtrack(cur, 0, 0)

	return res
}

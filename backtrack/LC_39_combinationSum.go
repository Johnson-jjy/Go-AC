package backtrack

import "sort"

// 组合问题: 1.本质上这是一道完全背包,但因为要输出组合情况,故建议用回溯,若是求数量,可按完全背包处理
// 2.只是组合不强求顺序则可对原数组进行排序,并做一个初步判断
// 3.注意剪枝
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	if target < candidates[0] {
		return res
	}
	cur := make([]int, 0)

	var backtrack func(cur []int, sum int, index int)
	backtrack = func(cur []int, sum int, index int) {
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			// 这个if判断也可以放在for上面,但放在这里可少走后面的for循环,因为数组是有序的
			if sum + candidates[i] > target {
				return
			}
			sum += candidates[i]
			backtrack(cur, sum, i)
			cur = cur[:len(cur) - 1]
			sum -= candidates[i]
		}
	}

	backtrack(cur, 0, 0)

	return res
}

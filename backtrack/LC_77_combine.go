package backtrack

// 求组和: 1.不重复->不用used判重; 2.求组和->暗含有序,调整startIndex;
// 3.考虑当前量和剩余量的关系,可以剪枝
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	var backtrack func(cur []int, index int)
	backtrack = func(cur []int, index int) {
		if len(cur) + (n - index + 1) < k {
			return
		}

		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := index; i <= n; i++ {
			cur = append(cur, i)
			backtrack(cur, i + 1)
			cur = cur[:len(cur) - 1]
		}
	}

	backtrack(cur, 1)

	return res
}

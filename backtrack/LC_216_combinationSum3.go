package backtrack

// 组合总和3: 限定了目标n和k ->注:可根据n和k进行剪枝
// 注意: 1~9的范围别弄错了
// 另待补充: 基于二进制子集的优化
// 算法注: backtrack必须先var再定义;因为其调用了自身的递归,而check则可以直接定义使用->没调用自身
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)

	// 注: 不要被题目条件的k==3迷惑,没有这个条件
	if n < 3 || n > 45 {
		return res
	}

	var backtrack func(cur []int, num int, sum int, index int)
	backtrack = func(cur []int, num int, sum int, index int) {
		if num == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, cur)
				res = append(res, tmp)
			}
			return
		}

		for i := index; i <= 9; i++ { // 可以等于9
			if sum + i > n || n - sum > 9 * (k - num) {
				return
			}
			cur = append(cur, i)
			sum += i
			backtrack(cur, num + 1, sum, i + 1)
			cur = cur[:len(cur) - 1]
			sum -= i
		}
	}

	cur := make([]int, 0)
	backtrack(cur, 0, 0, 1) // 无0

	return res
}

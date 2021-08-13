package backtrack

// 括号查找: 1.明确不能称为括号匹配的情况; 2.不想标准的for循环里走回溯,但本质类似二叉树,直接不用for,进行相关递归即可
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	cur := make([]byte, 0)

	var backtrack func(cur []byte, left int, right int)
	backtrack = func(cur []byte, left int, right int) {
		if left > n || right > n || right > left {
			return
		}

		if right == n {
			tmp := string(cur)
			res = append(res, tmp)
			return
		}

		cur = append(cur, '(')
		backtrack(cur, left + 1, right)
		cur = cur[:len(cur) - 1]
		cur = append(cur, ')')
		backtrack(cur, left, right + 1)
		cur = cur[:len(cur) - 1]
	}

	backtrack(cur, 0, 0)

	return res
}

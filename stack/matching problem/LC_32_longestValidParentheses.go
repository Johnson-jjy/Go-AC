package matching_problem

func longestValidParentheses(s string) int {
	stack := make([]int, 0)

	pre := -1
	res := 0
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) != 0 {
				stack = stack[:len(stack) - 1]
				if len(stack) == 0 {
					cur = pre
				} else {
					cur = stack[len(stack) - 1]
				}
				if res < i - cur {
					res = i - cur
				}
			} else {
				pre = i
			}
		}
	}

	return res
}
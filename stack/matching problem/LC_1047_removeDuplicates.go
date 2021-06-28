package matching_problem

func removeDuplicates(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			cur := stack[len(stack) - 1]
			if s[i] == cur {
				stack = stack[: len(stack) - 1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}

	//注：别写成 (string)stack
	return string(stack)
}

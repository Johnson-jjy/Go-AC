package monotonic_stack

// 经典标识--"下一个最高"--根据题目需求,保存坐标即可
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {
			cur := stack[len(stack) - 1]
			res[cur] = i - cur
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	return res
}

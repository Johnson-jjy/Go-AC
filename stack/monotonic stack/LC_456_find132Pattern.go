package monotonic_stack

func find132pattern(nums []int) bool {
	stack := make([]int, 1)
	flag := nums[0]
	stack[0] = nums[0]

	for i := 0; i < len(nums); i++ {
		flag = max(flag, stack[len(stack) - 1])
		for len(stack) > 0 && nums[i] < stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		}
		if nums[i] < flag && len(stack) > 0 {
			return true
		}
		stack = append(stack, nums[i])
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

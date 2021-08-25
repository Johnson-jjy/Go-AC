package monotonic_stack

// 下一个更大元素Ⅱ
// 找下一个最大的 -> 依然想到单调栈
// 循环数组 -> 可将数组范围扩大一倍 -> 辨析打家劫舍
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
			cur := stack[len(stack) - 1]
			res[cur] = nums[i]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
			cur := stack[len(stack) - 1]
			res[cur] = nums[i]
			stack = stack[:len(stack) - 1]
		}
	}
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		res[cur] = -1
		stack = stack[:len(stack) - 1]
	}

	return res
}

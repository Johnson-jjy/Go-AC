package monotonic_stack

// 柱状图中的最大矩形
// 1. 单调栈: 维持一个单调递增栈则元素弹出时:要加进来的i是其后第一个比它小的;弹出后栈顶元素(如果有的话)是其前第一个比它小的
// 2. 注意各个坐标 -> stack中保存的是heights的坐标, 但也要根据stack的情况,算需要取的stack坐标 -> 注意非空的判断
// 待补充: 哨兵技巧
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)

	res := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack) - 1]] > heights[i] {
			n := len(stack) - 1
			h := heights[stack[n]] // 注意: n只是stack的最后一个元素的坐标, 取其值才是对应heights的index
			left := -1
			if n > 0 {
				left = stack[n - 1]
			}
			cur := h * (i - left - 1)
			if cur > res {
				res = cur
			}
			stack = stack[:n]
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		n := len(stack) - 1
		h := heights[stack[n]]
		left := -1
		if n > 0 {
			left = stack[n - 1]
		}
		cur := h * (len(heights) - left - 1)
		if cur > res {
			res = cur
		}
		stack = stack[:n]
	}

	return res
}

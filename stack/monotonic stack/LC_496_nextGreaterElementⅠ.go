package monotonic_stack

// 形如找到"下一个最大/小的数(大小或坐标)" -> 联想单调栈! -> 根据题目需求,本题栈中直接保存值
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	store := make(map[int]int, len(nums1)) // 对于num1中相关数字的对应情况,用map保存
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack) - 1] {
			cur := stack[len(stack) - 1]
			store[cur] = nums2[i]
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, nums2[i])
	}
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		store[cur] = -1
		stack = stack[:len(stack) - 1]
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = store[nums1[i]]
	}

	return nums1
}

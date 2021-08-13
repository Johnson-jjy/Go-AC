package backtrack

// 求子集: 1.无重复元素->不涉及used去重 2.暗含顺序->startIndex
// 注: 每次startIndex应该由for循环中的i+1而定,而不是传入的index+1定
var res78 [][]int

func subsets(nums []int) [][]int {
	res78 = make([][]int, 0)
	cur := make([]int, 0)

	backtrack78(cur, nums, 0)

	return res78
}

func backtrack78(cur []int, nums []int, index int) {
	if index > len(nums) {
		return
	}
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	res78 = append(res78, tmp)

	for i := index; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrack78(cur, nums, i + 1) // 不是index+1
		cur = cur[:len(cur) - 1]
	}
}

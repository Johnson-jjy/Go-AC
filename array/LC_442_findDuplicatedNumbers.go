package array

// 数组中元素值的范围确定->原地hash: +n -> 超出一定范围则出现两次(也可以用负数)
func findDuplicates(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
		if nums[v] > n + n {
			res = append(res, v + 1)
		}
	}

	return res
}

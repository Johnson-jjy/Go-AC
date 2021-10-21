package two_pointers

// 有序数组的平方
// 思路: 双指针, 从两端(可能的大值)往中间走
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left := 0
	right := n - 1

	index := n - 1
	for left <= right {
		n1 := nums[left] * nums[left]
		n2 := nums[right] * nums[right]
		if n1 > n2 {
			res[index] = n1
			left++
		} else {
			res[index] = n2
			right--
		}
		index--
	}

	return res
}

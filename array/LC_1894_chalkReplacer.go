package array

// 找到需要补充粉笔的学生编号
// 思路: 前缀和 + 二分: 通过前缀和找到余下不足一轮的量;在根据其有序性找目标
// 注: 按题目要求,因为需要严格小于,故目标应该满足 "<"; 且left定为1;right定位n; 其中逻辑需细品
func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	preSum := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i - 1] + chalk[i - 1]
	}

	k %= preSum[n]

	left := 1
	right := n
	ans := left
	for left <= right {
		mid := left + (right - left)/2
		if preSum[mid] > k {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans - 1
}

package offer

// 注意ans初值的设定->应对缺失值恰好为n-1的情况
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	ans := len(nums)

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == mid {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}

	return ans
}

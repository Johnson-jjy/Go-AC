package offer

// 本解法采用两次二分,意在熟悉对于第一个==target的值和最后一个==target的值的查找
// 还可用一次二分+线性扫描(不再赘述了)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	first := -1
	// 找第一个等于target的值
	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			first = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if first == -1 {
		return 0
	}

	// 找最后一个等于target的值
	left = 0
	right = len(nums) - 1
	last := len(nums)
	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return last - first + 1
}

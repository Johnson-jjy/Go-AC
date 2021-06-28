package Binary_Search

// 当前解法不触及人心；还需进一步完善修改
func findMin(nums []int) int {
	flag := nums[0]
	left := 0
	m := len(nums) - 1
	right := m
	mid := 0

	if m == 0 || nums[0] < nums[m] {
		return nums[0]
	}
	if nums[m] < nums[m - 1] {
		return nums[m]
	}

	for left <= right {
		mid = left + (right - left)/2
		if nums[mid] < nums[mid + 1] && nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if nums[mid] >= flag {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

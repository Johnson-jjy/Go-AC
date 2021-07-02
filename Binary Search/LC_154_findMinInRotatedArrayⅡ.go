package Binary_Search

func findMin2(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right && nums[left] == nums[right] {
		right--
	}

	if left > right || nums[left] < nums[right] {
		return nums[left]
	}

	left++
	flag := left
	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if nums[mid] >= nums[flag] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[flag]
}

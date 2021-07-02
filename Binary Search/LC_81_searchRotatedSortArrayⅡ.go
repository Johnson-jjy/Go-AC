package Binary_Search

func search2(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right && nums[left] == nums[right] {
		right--
	}

	if nums[left] == target {
		return true
	}
	flag := left
	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			if target < nums[flag] && nums[mid] >= nums[flag] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[flag] && nums[mid] < nums[flag] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}
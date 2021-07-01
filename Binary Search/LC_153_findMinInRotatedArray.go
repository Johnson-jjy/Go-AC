package Binary_Search

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if nums[left] <= nums[right] {
		return nums[left]
	}

	// 注意此处对解空间的缩减
	left++
	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
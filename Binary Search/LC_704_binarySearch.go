package Binary_Search

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	mid := -1

	for left <= right {
		mid = left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

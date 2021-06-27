package Binary_Search

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums) - 1
	ans := -1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	res[0] = ans

	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	res[1] = ans

	return res
}

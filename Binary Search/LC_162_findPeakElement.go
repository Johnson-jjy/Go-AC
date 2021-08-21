package Binary_Search

// 寻找峰值
// 1.对首元素和尾元素进行判断; 注意返回index而不是返回值; -> 注:对尾index返回时,不要返回错了
// 2.核心: 因为-1和len可看作-∞; 故必然存在峰值; 也就是说,必然存在靠近头或者尾的单调区间; 于是可根据对mid处在哪段区间来判断峰值可能在哪边
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums) - 2] < nums[len(nums) - 1] {
		return len(nums) - 1
	}

	left := 1
	right := len(nums) - 2

	for left <= right {
		mid := left + (right - left)/2
		if nums[mid] > nums[mid + 1] && nums[mid] > nums[mid - 1] {
			return mid
		} else if nums[mid] > nums[mid + 1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

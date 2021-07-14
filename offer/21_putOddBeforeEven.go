package offer

// 解1 : 首尾双指针 + for与多判据
func exchange1(nums []int) []int {
	left := 0
	m := len(nums) - 1
	right := m

	for left < right {
		for left < right && left <= m && nums[left] % 2 == 1 {
			left++
		}
		for left < right && right >= 0 && nums[right] % 2 == 0 {
			right--
		}
		if left < right {
			temp := nums[right]
			nums[right] = nums[left]
			nums[left] = temp
		}
	}

	return nums
}

// 解2 : 快慢双指针 + if-continue
func exchange(nums []int) []int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] % 2 == 1 {
			temp := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = temp
			slow++
		}
		fast++ // 注:快指针总是会移动
	}

	return nums
}
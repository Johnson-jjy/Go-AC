package Binary_Search

func search2(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	/*
	这段代码为理解的重难点：
	1.当左右相等的时候，此时整个数组不具有二段性，无法二分
	2.可能整个数组会全部匹配，所以需要在中间过程加入对是否恰好为target的判断
	*/
	for left <= right && nums[left] == nums[right] {
		if nums[left] == target {
			return true
		}
		left++
		right--
	}

	// 缩减了原数组的区间后，再用flag标记此时的左端点
	flag := left

	// 以下逻辑已和LC_33相同
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

package Binary_Search

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
**/
type MountainArray struct {
}

// 伪实现
func (this *MountainArray) get(index int) int { return -1 }
func (this *MountainArray) length() int { return -1 }


func findInMountainArray(target int, mountainArr *MountainArray) int {
	m := mountainArr.length()
	left := 1
	right := m - 2
	first := mountainArr.get(0)
	last := mountainArr.get(m - 1)

	if target < first && target < last {
		return -1
	}

	peak := right
	for left <= right {
		mid := left + (right - left) / 2

		if mountainArr.get(mid) > mountainArr.get(mid - 1) {
			peak = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if target > mountainArr.get(peak) {
		return -1
	}

	left = 0
	right = peak
	for left <= right {
		mid := left + (right - left) / 2

		cur := mountainArr.get(mid)
		if cur == target {
			return mid
		} else if cur > target{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left = peak + 1
	right = m - 1
	for left <= right {
		mid := left + (right - left) / 2

		cur := mountainArr.get(mid)
		if cur == target {
			return mid
		} else if cur < target{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

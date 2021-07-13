package offer

// 经典二分题,通过本题吃透本质
func minArray(numbers []int) int {
	m := len(numbers)
	if m == 1 {
		return numbers[0]
	}
	left := 0
	right := m - 1
	for left <= right && numbers[left] == numbers[right] {
		right--
	}
	if left > right || numbers[left] < numbers[right] {
		return numbers[left]
	}

	left++
	flag := left
	for left <= right {
		mid := left + (right - left)/2
		if numbers[mid] < numbers[mid - 1] {
			return numbers[mid]
		} else if numbers[mid] >= numbers[flag] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return numbers[flag]
}

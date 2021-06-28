package Binary_Search

// 最接近原始二分的排除思路
func mySqrt1(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := left + (right - left)/2
		if mid * mid <= x && (mid + 1) * (mid + 1) > x {
			return mid
		} else if mid * mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// 用ans保存可能的答案
func mySqrt2(x int) int {
	left := 0
	right := x
	ans := 0

	for left <= right {
		mid := left + (right - left)/2
		if mid * mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
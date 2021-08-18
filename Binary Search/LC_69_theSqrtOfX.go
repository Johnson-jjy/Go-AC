package Binary_Search

// x的平方根
// 最接近原始二分的排除思路
func mySqrt1(x int) int {
	left := 0
	right := x // 注意，没必要 x/2; 还有x为1这种情况

	ans := 0
	for left <= right {
		mid := left + (right - left)/2
		if mid * mid == x {
			return mid
		} else if mid * mid < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
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
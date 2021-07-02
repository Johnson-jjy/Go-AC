package two_pointers

import "math"

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		cur := left * left + right * right
		if cur == c {
			return true
		} else if cur > c {
			right--
		} else {
			left++
		}
	}

	return false
}

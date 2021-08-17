package greedy

import "strconv"

// 单调递增的数字
// 从前往后不能保持无后效性,但从后往前->只有要需要改成9的,之后的不会影响
// 从后往前,找到第一个需要将后面都修改为9的即可
// 注: - '0'
func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}
	str := strconv.Itoa(n)
	s := []byte(str)
	index := len(s)

	for i := len(s) - 1; i >= 1; i-- {
		if s[i - 1] > s[i] {
			index = i
			s[i - 1] -= 1
		}
	}
	sum := 0
	for i := 0; i < index; i++ {
		sum = sum * 10 + int(s[i] - '0')
	}
	for i := index; i < len(s); i++ {
		sum = sum * 10 + 9
	}

	return sum
}

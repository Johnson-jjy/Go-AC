package math

import "strconv"

// 6和9组成的最大数字
func maximum69Number (num int) int {
	s := strconv.Itoa(num)
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == '6' {
			b[i] = '9'
			break
		}
	}

	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum * 10 + int(b[i] - '0')
	}

	return sum
}


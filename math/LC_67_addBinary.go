package math

import "strconv"

// 二进制求和
// 解一: 类似于字符串求和
// 解二: 位运算相关,待补充
func addBinary(a string, b string) string {
	res := ""
	left := len(a) - 1
	right := len(b) - 1
	carry := 0
	for left >= 0 || right >= 0 || carry > 0 {
		l := 0
		if left >= 0 {
			l = int(a[left] - '0')
			left--
		}
		r := 0
		if right >= 0 {
			r = int(b[right] - '0')
			right--
		}
		cur := l + r + carry
		carry = cur/2
		res = strconv.Itoa(cur%2) + res
	}

	return res
}

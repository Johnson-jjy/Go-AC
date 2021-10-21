package two_pointers

// 反转字符串Ⅱ

// 思路: 双指针往中走, 不断交换; 明确当前left和right所处的位置即可
func reverseStr(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i = i + k * 2 {
		reverse_(t, i, i + k - 1)
	}

	return string(t)
}

func reverse_(s []byte, start int, end int) {
	if end >= len(s) {
		end = len(s) - 1
	}
	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

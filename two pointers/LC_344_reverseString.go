package two_pointers

// 反转字符串
// 思路: 双指针往中间
func reverseString(s []byte)  {
	left := 0
	right := len(s) - 1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

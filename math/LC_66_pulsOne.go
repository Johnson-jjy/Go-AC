package math

// 加一
// 思路: 从后往前, 加一即可, 遇到9时考虑进位的情况; 注意一直进位的特殊情况 --> 需要头部补1
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		cur := digits[i] + carry
		if cur <= 9 {
			carry = 0
			digits[i] = cur
			break
		}
		digits[i] = 0
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

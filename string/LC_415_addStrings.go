package string

import "strconv"

// 经典题: 核心思想 -- 模拟 -> 从右往左扫,做相加,保存进位数字即可 -> 注: 用||
// 注: go相关strconv库的使用;类型转换;字符串的拼接
func addStrings(num1 string, num2 string) string {
	left := len(num1) - 1
	right := len(num2) - 1
	var res string

	carry := 0
	for left >= 0 || right >= 0{
		n1 := 0
		n2 := 0
		if left >= 0 {
			n1 = (int)(num1[left] - '0')
			left--
		}
		if right >= 0 {
			n2 = (int)(num2[right] - '0')
			right--
		}
		cur := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		res = strconv.Itoa(cur) + res
		//fmt.Println(cur, res)
	}
	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}

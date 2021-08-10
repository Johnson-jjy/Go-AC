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
			n1 = int(num1[left] - '0') // 注: 这里一定要转换成int, go无隐式转换
			left-- // 这里一定要--,否则进死循环
		}
		if right >= 0 {
			n2 = int(num2[right] - '0')
			right--
		}
		cur := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		res = strconv.Itoa(cur) + res
		//fmt.Println(cur, res)
	}
	if carry > 0 { // 这里是>0, 勿弄错
		res = strconv.Itoa(carry) + res
	}

	return res
}


// 解法2: 用byte数组存,最后逆序
func addStrings2(num1 string, num2 string) string {
	res := make([]byte, 0)

	carry := 0
	left := len(num1) - 1
	right := len(num2) - 1

	for left >= 0 || right >= 0 {
		l := 0
		r := 0
		if left >= 0 {
			l = int(num1[left] - '0')
			left--
		}
		if right >= 0 {
			r = int(num2[right] - '0')
			right--
		}
		cur := byte((l + r + carry)%10 + '0') // 注意,这里一定要+'0'
		carry = (l + r + carry)/10

		res = append(res, cur)
	}

	if carry > 0 {
		cur := byte(carry + '0') // 同理,这里也得+'0'
		res = append(res, cur)
	}

	n := len(res)
	for i := 0; i < n/2; i++ { // 注意,这里应该是<, 否则偶数时最中间的两个不会转换
		res[i], res[n - 1 - i] = res[n - 1 - i], res[i]
	}

	return string(res)
}
package main

import "fmt"

// 两个超大整数,最大不超过100万位求和;
// 本质即字符串加法,注意! res用int型数组保存即可,不需要再对cur值进行处理
func main() {
	var s1, s2 string
	s1 = "0"
	s2 = "0"

	carry := 0
	left := len(s1) - 1
	right := len(s2) - 1
	var res []int

	for left >= 0 || right >= 0 || carry > 0 {
		l := 0
		if left >= 0 {
			l = int(s1[left] - '0')
			left--
		}
		r := 0
		if right >= 0 {
			r = int(s2[right] - '0')
			right--
		}
		cur := l + r + carry
		res = append(res, cur % 10)
		carry = cur / 10
	}

	i := 0
	j := len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	for k := 0; k < len(res); k++ {
		fmt.Print(res[k])
	}
}

package math

import "math"

func myAtoi(s string) int {
	MAX := math.MaxInt32
	MIN := math.MinInt32
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}
	if index == len(s) {
		return 0
	}
	flag := false // t-负,f-正
	if s[index] == '-' {
		flag = true
		index++
	} else if  s[index] == '+' {
		index++
	} else if (s[index] >'9' || s[index] < '0') {
		return 0
	}
	// 若之前读到的不是数字则直接为0
	// for index < len(s) && (s[index] >'9' || s[index] < '0') {
	//     index++
	// }
	// if index == len(s) {
	//     return 0
	// }
	res := 0
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		cur := int(s[index] - '0')
		// if !flag && res >= MAX/10 {
		//     return MAX
		// }
		// if flag && -1 * res <= MIN/10 {
		//     return MIN
		// }
		res = res * 10 + cur
		if !flag && res >= MAX {
			return MAX
		}
		if flag && -1 * res <= MIN {
			return MIN
		}
		index++
	}
	// 返回int，故不需要考虑以下小数部分 -> 注：小数部分又将涉及大量类型转换
	// if index == len(s) || (s[index] != '.' && (s[index] > '9' || s[index] < '0')) {
	//     return res
	// }
	// index++

	// small := 0.1
	// var sp = float64(0)
	// for index < len(s) && s[index] >= '0' && s[index] <= '9' {
	//     sp += small * (float64(s[index] - '0'))
	//     small *= 0.1
	// }

	if flag { // 注意返回负值
		return -1 * res
	}
	return res
}

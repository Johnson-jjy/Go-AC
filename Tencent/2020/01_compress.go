package main

import "fmt"

//import "fmt"

/**
* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
*
*
* @param str string字符串
* @return string字符串
*/
func compress( str string ) string {
	// write code here
	res := make([]byte, 0)
	stack := make([]byte, 0)
	pre := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if len(stack) == 0 && str[i] != '[' {
			res = append(res, str[i])
		} else {
			if str[i] == ']' {
				tmp := make([]byte, len(pre))
				copy(tmp, pre)
				n := len(stack) - 1
				for stack[n] != '|' {
					tmp = append(tmp, stack[n])
					n--
				}
				n-- // 取数字
				num := int(stack[n] - '0')
				curStr := make([]byte, 0)
				for num > 0 {
					curStr = append(curStr, tmp...)
					num--
				}
				pre = make([]byte, len(curStr))
				copy(pre, curStr)
				n-- // 取左边界
				stack = stack[:n]
				if len(stack) == 0 {
					tmp := reverse(pre)
					res = append(res, tmp...)
					pre = make([]byte, 0)
				}
			} else {
				stack = append(stack, str[i])
			}
		}
	}
	//fmt.Println(res)
	return string(res)
}

func reverse(chars []byte) []byte {
	i := 0
	j := len(chars) - 1
	for i < j {
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}

	return chars
}

func main() {
	fmt.Println(compress("HG[3|B[2|CA][2|DE]]F")) // 拼接的话,就会出问题
}
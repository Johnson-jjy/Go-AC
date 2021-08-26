package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 输入一个带字符宽度控制的字符串和文本框宽度， 返回文本框内的最后一个可见字符
 * @param inputText string字符串 输入的字符串
 * @param limitSize int整型 文本框宽度
 * @return char字符型
 */
func GetLastVisibleCharInBox( inputText string ,  limitSize int ) byte {
	// write code here
	n := len(inputText)
	i := 0
	stack := make([]int, 0)
	flag := 16
	cur := 0
	pre := inputText[0]
	for i < n {
		if inputText[i] == '<' {
			i++
		} else if inputText[i] == '>' {
			i++
		} else if inputText[i] >= '0' && inputText[i] <= '9' && i > 1 && inputText[i - 1] == '<' {
			sum := 0
			j := i
			for inputText[j] >= '0' && inputText[j] <= '9'{
				sum = sum * 10 + int(inputText[j] - '0')
				j++
			}
			flag = sum
			stack = append(stack, sum)
			i = j
		} else if inputText[i] == '/' {
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				flag = 16
			} else {
				flag = stack[len(stack) - 1]
			}
			i++
		} else {
			cur += flag
			if cur > limitSize {
				break
			}
			pre = inputText[i]
			i++
		}
	}

	return pre
}
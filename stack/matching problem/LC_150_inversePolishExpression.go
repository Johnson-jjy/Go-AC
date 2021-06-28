package matching_problem

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		//掌握Atoi的用法！
		cur, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, cur)
		} else {
			num1 := stack[len(stack) - 1]
			num2 := stack[len(stack) - 2]
			//记得将栈顶的两个数字弹出去
			stack = stack[:len(stack) - 2]
			//注意作用域，需要在switch外声明
			curNum := 0
			//注意不是cur，cur已经被转换成了某数字
			switch tokens[i] {
			case "+":
				curNum = num2 + num1
			case "-":
				curNum = num2 - num1
			case "*":
				curNum = num2 * num1
			case "/":
				curNum = num2 / num1
			}
			stack = append(stack, curNum)
		}
	}

	return stack[0]
}

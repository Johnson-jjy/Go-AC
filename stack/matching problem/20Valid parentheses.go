package matching_problem

func isValid(s string) bool {
	m := len(s)

	//奇数直接报错
	if m % 2 != 0 {
		return false
	}

	//建议：栈的初始化都用0(我一开始用的m/2)，否则，对其内元素操作后，底层数组会有默认值，最后不好判断
	stack := make([]byte, 0)

	//思路：左符号压入栈中，右符号则取栈顶元素并判断
	for i := 0; i < m; i++ {
		//注1：这里可以用一个map存三种匹配情况，类似于字典值对应，就可以不写这些冗长的if-else语句了
		//注2：对栈（以及队列）进行操作时，往往要判空
		//注3：流程控制语句的格式一定要注意
		if s[i] == ')' {
			if len(stack) == 0 || stack[len(stack) - 1] != '(' {
				return false
			}
			stack = stack[ : len(stack) - 1]
		} else if s[i] == '}' {
			if len(stack) == 0 || stack[len(stack) - 1] != '{' {
				return false
			}
			stack = stack[ : len(stack) - 1]
		} else if s[i] == ']' {
			if len(stack) == 0 || stack[len(stack) - 1] != '[' {
				return false
			}
			stack = stack[ : len(stack) - 1]
		} else {
			stack = append(stack, s[i])
			//fmt.Printf("%d\t%d\n", i, len(stack)) 此处用于检验
		}
	}

	//注1：返回语句可更间接地写为 return len(stack) == 0
	//注2：go中无三元操作符，需要用流程控制代替
	if len(stack) != 0 {
		return false
	}

	return true
}

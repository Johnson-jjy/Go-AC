package monotonic_stack

// 移掉K位数字
// 思路:做法无非就两种,贪心/DP -> 要移除k个数字，那显然得设计一个算法一个一个依次移除(如果要考虑一次移除多个数字的话，那复杂度应该就是用组合数公式算了,不切实际)
// 故贪心 -> 从前往后遍历这串数，遇到第一个下降的数字的时候移除前一个数字。这样可以保证这串数下降的大小最大
// 对数据的具体存储方式可用单调栈
// 细节: 1.字串首字符不能为0,故首0可直接不入栈; 2.若逆序的数字剪完后,还需要减少正序靠后的数字; 3.若最后为空字串,需要返回"0"
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}

	store := []byte(num)
	stack := make([]byte, 0)
	// 单调递增栈处理可能逆序的元素
	for i := 0; i < len(store); i++ {
		// 注意是for循环不是if
		for len(stack) > 0 && store[i] < stack[len(stack) - 1] && k > 0 {
			stack = stack[:len(stack) - 1]
			k--
		}
		if len(stack) == 0 && store[i] == '0' { // 对不合法的0不予考虑
			continue
		}
		stack = append(stack, store[i])
		//fmt.Println(stack)
	}
	//fmt.Println(stack)
	//然后再处理没处理完的k
	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack) - 1]
		k--
	}
	// 若最后栈空则可直接返回
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

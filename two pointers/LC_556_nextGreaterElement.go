package two_pointers

import "math"

// 下一个更大元素->类似下一个更大数: 1.注意在使用append的情况下,如何去调节之后的顺序
// 2.注意题目条件中,对于越界数的返回-1处理
// 另:可使用strconv的Itoa和Atoi语句,将其完全转变为LC_ 31来做
func nextGreaterElement(n int) int {
	if n <= 9 {
		return -1
	}
	flag := 0
	store := make([]int, 1)
	store[0] = n%10
	n /= 10
	for n > 0 {
		cur := n % 10
		store = append(store, cur)
		n /= 10
		if cur < store[len(store) - 2] {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return -1
	}

	for i := 0; i < len(store); i++ {
		if store[i] > store[len(store) - 1] {
			store[i], store[len(store) - 1] = store[len(store) - 1], store[i]
			break
		}
	}
	n = n * 10
	sum := store[len(store) - 1]
	for i := 0; i < len(store) - 1; i++ {
		n *= 10
		sum = sum * 10 + store[i]
	}
	if n + sum > math.MaxInt32 {
		return -1
	}
	return n + sum
}

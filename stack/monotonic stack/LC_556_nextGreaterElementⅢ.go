package monotonic_stack

import "math"

// 下一个更大元素Ⅲ
// 本题实际上不是单调栈类型的应用,但考虑到含义相近,故放在这里
// 从后往前找到第一组逆序对,然后再找这段区间中,小于index处值的数字,调换后,取最小即可
// 本题可讨巧,先将int转为string,然后变换为常规处理,也可就按int操作,但需要注意append是加在末尾的,故再调换操作时有些操作没必要
func nextGreaterElement3(n int) int {
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

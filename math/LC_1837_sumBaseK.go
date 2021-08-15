package math

func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		cur := n % k
		n /= k
		res += cur
	}

	return res
}

package string

//
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	store := make([]int, m + n)

	for i := m - 1; i >= 0; i-- {
		l := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			r := int(num2[j] - '0')
			cur := l * r
			store[i + j + 1] += cur
		}
	}

	res := make([]byte, 0)
	for i := m + n - 2; i >= 0; i-- {
		store[i] += store[i + 1]/10
		store[i + 1] %= 10
		res = append(res, byte(store[i + 1] + '0'))
	}
	if store[0] > 0 {
		res = append(res, byte(store[0] + '0'))
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res) - i - 1] = res[len(res) - i - 1], res[i]
	}

	return string(res)
}

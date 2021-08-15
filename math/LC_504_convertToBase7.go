package math

// 七进制数: 本题为NC中进制转换的一个特殊实例
// 注: 1.对于0要特判; 2.对于负数要转换符号
// 核心: 转换为进制是%; 每一轮迭代是/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := false
	if num < 0 {
		sign = true
		num = -num
	}
	res := ""
	for num > 0 {
		cur := num%7
		var b byte
		b = byte(cur + '0')
		res = string(b) + res
		num /= 7
	}

	if sign {
		return "-" + res
	}

	return res
}

package math

// 数字的补数
// 解一: do-while 形式求得num位数的每一位均为1时的数字, 相减即得结果
func findComplement1(num int) int {
	sum := 0
	count := 1
	origin := num
	for {
		num /= 2
		sum += count
		count *= 2
		if num == 0 {
			break
		}
	}

	return sum - origin
}

// 解二: 纯位运算操作
// 先找到最高位; 进而求二进制的反; 最后异或 --> 三个操作都可以是更高阶的位运算题目的基础
func findComplement2(num int) int {
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}

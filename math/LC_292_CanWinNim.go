package math

// Nim游戏
// 走到4的倍数必输, 因为可以控制二者和为4的倍数
func canWinNim(n int) bool {
	if n % 4 == 0 {
		return false
	}

	return true
}

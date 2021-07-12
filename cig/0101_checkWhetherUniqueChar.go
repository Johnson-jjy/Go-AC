package cig

// 哈希思想,已知字符数量则直接用切片
func isUnique(astr string) bool {
	store := make([]int, 26)
	for i := 0; i < len(astr); i++ {
		if store[astr[i] - 'a'] > 0 {
			return false
		}
		store[astr[i] - 'a']++
	}
	return true
}

// 位运算:待补充
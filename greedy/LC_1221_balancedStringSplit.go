package greedy

// 分割平衡字符串: 归纳法证明即可
func balancedStringSplit(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		}  else {
			count--
		}
		if count == 0 {
			res++
		}
	}

	return res
}

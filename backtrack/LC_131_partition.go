package backtrack

// 分割回文串: 1.判断回文串: -> 双指针即可
// 2.子集枚举（切割）-> 回溯 -> 注意index
func partition(s string) [][]string {
	res := make([][]string, 0)
	store := make(map[string]bool, 0)

	check := func(t string) bool {
		i := 0
		j := len(t) - 1
		for i < j {
			if t[i] == t[j] {
				i++
				j--
			} else {
				return false
			}
		}
		store[t] = true
		return true
	}

	var backtrack func(cur []string, index int)
	backtrack = func(cur []string, index int) {
		if index == len(s) {
			tmp := make([]string, len(cur))
			for i := 0; i < len(cur); i++ {
				tmp[i] = cur[i]
			}
			res = append(res, tmp)
			return
		}

		for i := index; i < len(s); i++ {
			tmp := s[index : i + 1]
			if store[tmp] || check(tmp) {
				cur = append(cur, tmp)
				backtrack(cur, i + 1)
				cur = cur[:len(cur) - 1]
			}
		}
	}

	cur := make([]string, 0)
	backtrack(cur, 0)

	return res
}



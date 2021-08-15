package Sort

// 有效的字母异位词
// 本质思想 -- hash -> 注意 -'a'而不是 -'0'
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	store := make([]int, 26)
	for i := 0; i < len(s); i++ {
		store[int(s[i] - 'a')]++
	}
	for i := 0; i < len(t); i++ {
		cur := int(t[i] - 'a')
		store[cur]--
		if store[cur] < 0 {
			return false
		}
	}

	return true
}

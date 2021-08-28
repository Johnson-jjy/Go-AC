package offer

// 第一个只出现一次的字符
// 哈希表存频数
func firstUniqChar(s string) byte {
	store := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cur := int(s[i] - 'a')
		store[cur]++
	}
	for i := 0; i < len(s); i++ {
		cur := int(s[i] - 'a')
		if store[cur] == 1 {
			return s[i]
		}
	}
	return ' '
}

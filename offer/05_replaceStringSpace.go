package offer

// 用byte不用rune,注意一次追加...
func replaceSpace(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

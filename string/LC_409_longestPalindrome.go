package string

// 最长回文串
// 注:1.可以有一个字母只出现1次; 2.返回结果不要忘了偶数乘2
func longestPalindrome(s string) int {
	store := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		store[s[i]]++
	}
	res := 0
	odd := 0
	for _, v := range store {
		if odd < 1 && v % 2 == 1 {
			odd = 1
		}
		res += v/2
	}

	return res * 2 + odd
}

package cig

// 注意题目没说一定是小写字母,故还是需要开map
func canPermutePalindrome(s string) bool {
	odd := 0
	store := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		store[s[i]]++
		if store[s[i]] % 2 == 1 {
			odd++
		} else {
			odd--
		}
	}

	if odd == 0 || odd == 1 {
		return true
	}
	return false
}

// 位运算:待补充
package cig

func CheckPermutation(s1 string, s2 string) bool {
	store := make([]int, 26)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		store[s1[i] - 'a']++
	}
	for i := 0; i < len(s2); i++ {
		store[s2[i] - 'a']--
		if store[s2[i] - 'a'] < 0 {
			return false
		}
	}
	return true
}

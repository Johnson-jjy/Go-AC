package two_pointers

// 思路: 双指针滑动 -> 注意指针滑动和数字增减的时间
// 待补充: 直接存左指针的下标
func lengthOfLongestSubstring(s string) int {
	store := make(map[byte]int, 0)
	res := 0

	left := 0
	right := 0
	for right < len(s) {
		store[s[right]]++
		if v := store[s[right]]; v > 1 {
			res = max(res, right - left)
			//fmt.Println(res, left, right)
			for store[s[right]] > 1 {
				store[s[left]]--
				left++
			}
		}
		right++
	}
	res = max(res, right - left)

	return res
}


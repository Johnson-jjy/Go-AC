package offer

// 双指针->滑动窗口 + 哈下表
// 注意res初值
func lengthOfLongestSubstring(s string) int {
	slow, fast, res := 0, 0, 0
	store := make(map[byte]int, len(s))

	for fast < len(s) {
		store[s[fast]]++
		if store[s[fast]] > 1 {
			cur := fast - slow
			if res < cur {
				res = cur
			}
		}
		for store[s[fast]] > 1 {
			store[s[slow]]--
			slow++
		}
		fast++
	}
	if res < fast - slow { // res因为全部不重复而一直没更新
		res = fast - slow
	}
	return res
}

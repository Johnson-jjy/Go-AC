package Binary_Tree

// 反转字符串中的单词Ⅲ

// 思路: 双指针, 关键在于确定需要交换的节点
func reverseWords(s string) string {
	start := 0
	for s[start] == ' ' && start < len(s) {
		start++
	}
	end := start
	t := []byte(s)

	for end < len(t) {
		for end < len(t) && t[end] != ' ' {
			end++
		}
		reverse(t, start, end - 1)
		for end < len(t) && t[end] == ' ' {
			end++
		}
		start = end
	}

	return string(t)
}

func reverse(s []byte, start int, end int) {
	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

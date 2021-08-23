package two_pointers

import "math"

// 最小覆盖字串: 对于求一个大字符串中的一段小连续字串,我们常可以用滑动窗口计数
// 对于滑动窗口的使用: 分清楚左右指针指向的坐标是已经被计入的还是未被计入的
// 难点1: 对于原始需要匹配的字符的统计; right拓展时match--和left拓展时match++应该是完全对应的
// 难点2: 每次计数,都应该是满足条件的情况下,一定left前进行计数更新 -> 因为移动了left之后,可能就不满足条件了
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	check := make(map[byte]int, 0)
	need := 0
	for i := 0; i < len(t); i++ {
		if _, ok := check[t[i]]; !ok {
			check[t[i]] = 1
			need++
		} else {
			check[t[i]]++
		}
	}

	left := 0
	right := 0
	res := ""
	curLen := math.MaxInt32
	match := 0

	for right < len(s) {
		for right < len(s) && match < need {
			check[s[right]]--
			if check[s[right]] == 0 {
				match++
			}
			right++
		}

		for left < right && match >= need {
			if right - left < curLen {
				curLen = right - left
				res = s[left : right]
			}
			if _, ok := check[s[left]]; ok {
				check[s[left]]++
				if check[s[left]] == 1 {
					match--
				}
				//match--
			}
			left++
		}
	}

	if curLen == math.MaxInt32 {
		return ""
	}

	return res
}

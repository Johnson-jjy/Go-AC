package backtrack

// 电话号码的字母组合: 虽有组合,但每个小字母选自不同的集合,故本质类似于排列
// 注: 因为每次都是从一个新集合的第一位开始扫描,故不需要用used数组去重
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	// used := make([][]bool, len(digits))
	// for i := 0; i < len(digits); i++ {
	//     used[i] = make([]bool, 4)
	// }

	var backtrack func(cur []byte, index int)
	backtrack = func(cur []byte, index int) {
		if index == len(digits) {
			tmp := string(cur)
			res = append(res, tmp)
			return
		}

		s := getString(digits[index])
		for i := 0; i < len(s); i++ {
			cur = append(cur, s[i])
			// used[index][i] = true
			backtrack(cur, index + 1)
			cur = cur[:len(cur) - 1]
			// used[index][i] = false
		}
	}

	cur := make([]byte, 0)
	backtrack(cur, 0)

	return res
}

func getString(b byte) string {
	var s string
	switch b {
	case '2' : s = "abc"
	case '3' : s = "def"
	case '4' : s = "ghi"
	case '5' : s = "jkl"
	case '6' : s = "mno"
	case '7' : s = "pqrs"
	case '8' : s = "tuv"
	case '9' : s = "wxyz"
	}
	return s
}

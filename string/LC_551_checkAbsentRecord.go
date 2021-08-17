package string

// 学生出勤记录Ⅰ: 理解清楚题意后,一遍扫描即可
func checkRecord(s string) bool {
	count := 0
	continuous := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			count++
			continuous = 0
			if count >= 2 {
				return false
			}
		} else if s[i] == 'L' {
			continuous++
			if continuous >= 3 {
				return false
			}
		} else {
			continuous = 0
		}
	}

	return true
}

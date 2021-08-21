package string

// 比较版本号: 解析出每一小段; 然后求出该段数字, 再进行比较即可
// 注: 1.注意for循环不要越界; 2.题目条件又说均可以保存为数字,故可直接转换
// 另: 偷鸡不成蚀把米, 不如直接写一个函数, 每次将字符串转为相应的数字
func compareVersion(version1 string, version2 string) int {
	left := 0
	right := 0

	for left < len(version1) && right < len(version2) {
		start1 := left
		start2 := right
		for left < len(version1) && version1[left] != '.' {
			left++
		}
		s1 := version1[start1 : left]
		for right < len(version2) && version2[right] != '.' {
			right++
		}
		s2 := version2[start2 : right]
		cur1, cur2 := 0, 0
		index1, index2 := 0, 0
		for index1 < len(s1) {
			cur1 = cur1 * 10 + int(s1[index1] - '0')
			index1++
		}

		for index2 < len(s2) {
			cur2 = cur2 * 10 + int(s2[index2] - '0')
			index2++
		}

		if cur1 == cur2 {
			left++
			right++
		} else if cur1 > cur2 {
			return 1
		} else {
			return -1
		}
	}

	for left < len(version1) {
		start1 := left
		for left < len(version1) && version1[left] != '.' {
			left++
		}
		s1 := version1[start1 : left]
		cur1 := 0
		index := 0
		for index < len(s1) {
			cur1 = cur1 * 10 + int(s1[index] - '0')
			index++
		}
		if cur1 > 0 {
			return 1
		}
		left++
	}

	for right < len(version2) {
		start2 := right
		for right < len(version2) && version2[right] != '.' {
			right++
		}
		s2 := version2[start2 : right]
		cur2 := 0
		index := 0
		for index < len(s2) {
			cur2 = cur2 * 10 + int(s2[index] - '0')
			index++
		}
		if cur2 > 0 {
			return -1
		}
		right++
	}

	return 0
}

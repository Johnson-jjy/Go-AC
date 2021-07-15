package cig

// 双指针: 注意什么时候移动指针
func oneEditAway(first string, second string) bool {
	m := len(first)
	n := len(second)
	if m - n > 1 || n - m > 1 {
		return false
	}

	if m < n {
		temp := second
		second = first
		first = temp
	}
	i := 0
	j := 0
	test := 0
	for i < len(first) && j < len(second) {
		if first[i] == second[j] {
			i++
			j++
		} else {
			if m != n {
				i++
			} else {
				i++
				j++
			}
			test++
			//fmt.Printf("%d\n", test)
			if test > 1 {
				return false
			}
		}
	}

	return true
}

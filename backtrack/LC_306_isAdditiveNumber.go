package backtrack

// 累加数: 类似842->拆分成斐波那契,本题只需计数,更简洁
func isAdditiveNumber(num string) bool {
	getNum := func(start int, end int) int {
		if num[start] == '0' && start != end {
			return -1
		}
		sum := 0
		for start <= end {
			cur := int(num[start] - '0')
			sum = sum * 10 + cur
			start++
		}

		return sum
	}

	var backtrack func(cur []int, index int) bool
	backtrack = func(cur []int, index int) bool{
		if index == len(num) && len(cur) > 2{
			return true
		}

		for i := index; i < len(num); i++ {
			n := getNum(index, i)
			if n < 0 {
				break
			}
			if len(cur) < 2 || n == cur[len(cur) - 1] + cur[len(cur) - 2]{
				cur = append(cur, n)
				if backtrack(cur, i + 1) {
					return true
				}
				cur = cur[:len(cur) - 1]
			}
		}

		return false
	}

	cur := make([]int, 0)
	if backtrack(cur, 0) {
		return true
	}

	return false
}

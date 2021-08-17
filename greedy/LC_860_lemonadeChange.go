package greedy

// 柠檬水找零: 贪心 -> 找20时,先给出10
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		} else if bills[i] == 20 {
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five = five - 3
			} else {
				return false
			}
		}
	}

	return true
}

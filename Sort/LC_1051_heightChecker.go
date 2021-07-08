package Sort

// 计数排序
func heightChecker(heights []int) int {
	store := make([]int, 101)

	for i := 0; i < len(heights); i++ {
		store[heights[i]]++
	}

	sortIndex := 0
	ans := 0
	for i := 0; i <= 100; i++ {
		for store[i] > 0 {
			if heights[sortIndex] != i {
				ans++
			}
			sortIndex++
			store[i]--
		}
	}

	return ans
}

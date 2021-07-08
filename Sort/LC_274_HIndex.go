package Sort

// 计数排序
func hIndex(citations []int) int {
	n := len(citations)
	store := make([]int, n + 1)
	for i := 0; i < n; i++ {
		if citations[i] >= n {
			store[n]++
		} else {
			store[citations[i]]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i == n - 1 {
			citations[i] = store[n]
		} else {
			citations[i] = store[i + 1] + citations[i + 1]
		}
		if citations[i] >= i + 1{
			return i + 1
		}
	}

	return 0
}

package greedy

// 分发糖果: 两次遍历 -> 一次从左往右,一次从右往左: 注,这种某元素需要和前后元素进行对比的,往往都需要两次按不同顺序的遍历
// 注: 关于res的统计可在第二次从右往左的遍历中进行
// 另待补充: O(1)的方法
func candy(ratings []int) int {
	store := make([]int, len(ratings))
	store[0] = 1
	store[len(store) - 1] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i - 1] {
			store[i] = store[i - 1] + 1
		} else {
			store[i] = 1
		}
	}

	res := store[len(store) - 1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			store[i] = max(store[i], store[i + 1] + 1)
		}
		res += store[i]
	}

	//res := 0
	// for i := 0; i < len(store); i++ {
	//     res += store[i]
	// }

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package graph

// 思路: 哈希表存相邻元素 -> 巧用切片而不是用多个哈希表去存
// 注意: 向最终的结果数组添加元素时, 一定要明确究竟是添加K还是V
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	store := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		cur1 := adjacentPairs[i][0]
		cur2 := adjacentPairs[i][1]
		store[cur1] = append(store[cur1], cur2)
		store[cur2] = append(store[cur2], cur1)
	}
	//index := 0
	res := make([]int, n + 1)
	for k, v := range store {
		//fmt.Printf("1:%v\n",v)
		if len(v) == 1 {
			res[0] = k
			res[1] = v[0]
			break
		}
	}
	//fmt.Printf("2:%v\n", res)
	for i := 2; i <= n; i++ {
		cur := store[res[i - 1]]
		//fmt.Printf("3:%v\n", cur)
		if res[i - 2] == cur[0] {
			res[i] = cur[1]
		} else {
			res[i] = cur[0]
		}
	}

	return res
}

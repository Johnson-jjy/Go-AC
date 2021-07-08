package Sort

// 计数排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	store := make([]int, 1001) // 题目给了数组中元素的取值范围，也可以遍历数组，找最大值和最小值
	for i := 0; i < len(arr1); i++ {
		store[arr1[i]]++
	}

	sortIndex := 0 // 本解法采用了直接修改原arr1数组，也可以新开数组用append进行修改
	for i := 0; i < len(arr2); i++ {
		for store[arr2[i]] > 0 {
			arr1[sortIndex] = arr2[i]
			sortIndex++
			store[arr2[i]]--
		}
	}

	for i := 0; i < len(store); i++ {
		for store[i] > 0 {
			arr1[sortIndex] = i
			sortIndex++
			store[i]--
		}
	}

	return arr1
}

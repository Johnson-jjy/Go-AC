package main

import "fmt"

/*
5
3 2 4 4 5
4 3 5 2 1
*/

// 思路: 前缀和记录,然后从前向后删除 -> 该方法超时qaq
func main()  {
	var n int
	fmt.Scan(&n)
	weight := make([]int, n)
	//index := make([]int, n)
	preSum := make([]int, n + 1)
	for i := 0; i < n; i++ {
		fmt.Scan(&weight[i])
		preSum[i + 1] = preSum[i] + weight[i]
	}
	var index int
	for j := 0; j < n; j++ {
		fmt.Scan(&index)
		left := preSum[index] - weight[index - 1]
		right := preSum[n] - preSum[index]
		need := weight[index - 1]
		weight[index - 1] = 0
		for index <= n {
			preSum[index] -= need
			index++
		}
		cur := max(left, right)
		fmt.Println(cur)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

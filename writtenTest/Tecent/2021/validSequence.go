package main

import "fmt"

func main()  {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	res := 0
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n + 1)
	}

	for i := n - 1; i >= 1; i-- {
		dp[i][i] = arr[i - 1]
		for j := i + 1; j <= n; j++ {
			if j == i + 1 {
				dp[i][j] = min(arr[i - 1], arr[j - 1])
				res++
			} else {
				if arr[i - 1] <= dp[i + 1][j - 1] && arr[j - 1] <= dp[i + 1][j - 1] {
					res++
				}
				dp[i][j] = min(min(arr[i - 1], arr[j - 1]), dp[i + 1][j - 1])
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
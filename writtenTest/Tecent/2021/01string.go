package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a string
	fmt.Scan(&a)

	dp := make([][]int, n + 1)
	if a[0] == '0' {
		dp[1] = []int{0, 0}
	} else if a[0] == '1' {
		dp[1] = []int{1, 1}
	}

	ans := 0
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, 2)
		for j := 1; j < i; j++ {
			if a[j - 1] == a[i - 1] {
				pre := dp[j][1]
				if pre + 1 + dp[j][0] > dp[i][0] {
					dp[i][1] = pre + 1
					dp[i][0] = dp[j][0] + dp[i][1]
				}
			} else {
				if 1 + dp[j][0] > dp[i][0] {
					dp[i][1] = 1
					dp[i][0] = dp[j][0] + 1
				}
			}
		}
		//fmt.Println(i, dp[i][0], dp[i][1])
		if ans < dp[i][0] {
			ans = dp[i][0]
		}
	}
	fmt.Println(ans)
}

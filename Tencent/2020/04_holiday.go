package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	exercise := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&exercise[i])
	}
	fmt.Println(exercise)
	gym := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&gym[i]) // 注意这些细节啊! 输入输出好几次问题了!
	}
	fmt.Println(gym)

	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n + 1)
	}

	dp[0][1] = 1
	if exercise[0] == 0 {
		dp[1][1] = math.MaxInt32
	}
	if gym[0] == 0 {
		dp[2][1] = math.MaxInt32
	}
	for i := 2; i <= n; i++ {
		cur0 := dp[0][i - 1]
		if exercise[i - 2] == 1 {
			cur0 = min(cur0, dp[1][i - 1])
		}
		if gym[i - 2] == 1 {
			cur0 = min(cur0, dp[2][i - 1])
		}
		dp[0][i] = cur0 + 1

		dp[1][i] = math.MaxInt32
		if exercise[i - 1] == 1 {
			cur1 := dp[0][i - 1]
			//if gym[i - 2] == 1 {
			//	cur1 = min(cur1, dp[2][i - 1])
			//}
			cur1 = min(cur1, dp[2][i - 1])
			dp[1][i] = cur1
		}

		dp[2][i] = math.MaxInt32
		if gym[i - 1] == 1 {
			cur2 := dp[0][i - 1]
			//if exercise[i - 1] == 1 {
			//	cur2 = min(cur2, dp[1][i - 1])
			//}
			cur2 = min(cur2, dp[1][i - 1])
			dp[2][i] = cur2
		}
	}

	res := min(dp[0][n], dp[1][n])
	res = min(res, dp[2][n])

	fmt.Println(dp[0][n], dp[1][n], dp[2][n], res)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

package offer

import "math"

// dp: 注意转移方程的精妙实现->抓住了每一个丑数的得来规律
// 注意对相同乘积的消除->使用if而非if-else
// 本解法的min采用了可变参数,在LC的Oj上提交时,不采用可变参数会更快
func nthUglyNumber(n int) int {
	dp := make([]int, n + 1)
	dp[1] = 1
	p2 := 1
	p3 := 1
	p5 := 1
	for i := 2; i <= n; i++ {
		cur2 := dp[p2] * 2
		cur3 := dp[p3] * 3
		cur5 := dp[p5] * 5
		dp[i] = min(cur2, cur3, cur5)
		if dp[i] == cur2 {
			p2++
		}
		if dp[i] == cur3 {
			p3++
		}
		if dp[i] == cur5 {
			p5++
		}
	}
	return dp[n]
}

func min(args ...int) int {
	res := math.MaxInt32
	for _, v := range args {
		if v < res {
			res = v
		}
	}

	return res
}

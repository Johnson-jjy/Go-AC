package array

// 所有奇数长度子数组的和
// 解一: On的优解:直接计数每个数字需要加和的次数
func sumOddLengthSubarrays(arr []int) (ans int) {
	l := len(arr)
	p := (l + 1) / 2
	for i, v := range arr {
		k := 0
		k1 := (l - i + 1) / 2
		k += k1*(i+k1-l) + p*l
		k2 := i/2 + 1
		k -= k2*(i-k2+1) - p*(p+1)
		ans += k * v
	}
	return
}

// 前缀和暴力硬搜: On2
func sumOddLengthSubarrays2(arr []int) int {
	n := len(arr)
	preSum := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i - 1] + arr[i  - 1]
	}

	res := preSum[n]

	for i := 3; i <= n; i += 2 {
		for j := i; j <= n; j++ {
			res += preSum[j] - preSum[j - i]
		}
	}

	return res
}

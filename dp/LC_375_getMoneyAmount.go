package dp

// 猜数字大小Ⅱ
// 解:

// 二分?错误思路! -> 事实上,无论是用直接二分还是前缀和,都没有对所有可能出现的情况枚举完
func getMoneyAmount(n int) int {
	preSum := make([]int, n + 1)
	//store := make(map[int]int, n)
	for  i := 1; i <= n; i++ {
		preSum[i] = preSum[i - 1] + i
		//store[preSum[i]] = i
	}

	left := 1
	right := n
	sum := 0

	var findFirst func(target int) int
	findFirst = func(target int) int {
		left := 0
		right := len(preSum) - 1
		ans := right
		for left <= right {
			mid := left + (right - left)/2
			if preSum[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
				ans = mid
			}
		}

		return ans
	}

	//last := 0
	for left <= right - 2 {
		mid := preSum[left] + (preSum[right] - preSum[left] + 1)/2
		left = findFirst(mid)
		// fmt.Println(left, mid)
		sum += left
		//last = left
	}

	return sum
}

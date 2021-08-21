package subsequence

// 最长递增子序列
// 法一: 基础dp -> dp数组的含义: 以 i 为结尾的最长子序列的长度 -> 两层for里用 j < i 比较
// 法二: 二分加快 -> 见后文 -> 用一个数组保存当前每一个子序列最值的大小;对于当前值更大,则store++;或者更新该值的位置
func lengthOfLIS1(nums []int) int {
	m := len(nums)
	dp := make([]int, m + 1)
	dp[1] = 1
	res := 1
	for i := 2; i <= m; i++ {
		dp[i] = 1 // 注: 无论如何,都可以只取自己得1
		for j := 1; j < i; j++ { // 注意遍历顺序,j < i -> 每一个i的位置定了之后,比较在他之前的数字
			if nums[i - 1] > nums[j - 1] {
				dp[i] = max(dp[i], dp[j] + 1)
				if dp[i] > res {
					res = dp[i]
				}
			}
		}
	}

	return res
}

func lengthOfLIS2(nums []int) int {
	m := len(nums)
	store := make([]int, 1)
	store[0] = nums[0]

	for i := 1; i < m; i++ {
		n := len(store) - 1
		if nums[i] > store[n] {
			store = append(store, nums[i])
		} else {
			left := 0
			right := n
			ans := 0 // 注: ans记录store数组中第一个比nums[i]更大的值,然后将其替换
			for left <= right {
				mid := left + (right - left)/2
				if store[mid] >= nums[i] {
					right = mid - 1
					ans = mid
				} else {
					left = mid + 1
				}
			}
			store[ans] = nums[i]
		}
	}

	return len(store) // 最后返回store数组的长度即可
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

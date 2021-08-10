package math

// 遍历找满足等差数列的子序列 -> 对每个子序列计数即可
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	res := 0
	for i := 1; i <= len(nums) - 2; i++ {
		diff := nums[i] - nums[i - 1]
		n := 2
		for i <= len(nums) - 2 && diff == nums[i + 1] - nums[i] {
			n++
			i++
		}
		if n >= 3 {
			res += count(n)
		}
	}

	return res
}

func count(length int) int {
	m := length - 2

	return ((1 + m) * m) / 2
}

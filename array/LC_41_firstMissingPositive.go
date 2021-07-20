package array

// 解法一:原地hash -> 负数变正,小于n则将位置上的数变负,最后扫描确定first
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num - 1] = -abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 解法2:数字交换 -> 注意首位可能刚好是n的特殊情况
func firstMissingPositive2(nums []int) int {
	i := 0
	for i < len(nums) {
		temp := nums[i]
		if temp < 0 || temp >= len(nums) {
			i++
			continue
		} else {
			for temp >= 0 && temp < len(nums) && temp != nums[temp] {
				nums[i] = nums[temp]
				nums[temp] = temp
				temp = nums[i]
			}
			i++
		}
	}

	for j := 1; j < len(nums); j++ {
		if nums[j] != j {
			return j
		}
	}
	if nums[0] != len(nums) {
		return len(nums)
	}

	return len(nums) + 1
}



package array

// 解法一:原地hash
// 1.负数变正, 0也要变! -> 本质上是数组中没有意义的index就都要变 n + 1
// 2.小于n则将位置上的数变负 -> 注意变负必须套abs, 不能直接改符号,不然可能负负得正
// 3.最后扫描确定first -> index + 1
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 { // 负数全部先变为不处理的数; 同时对于0也要处理
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n { // 值为n时依然有index可处理
			// 不能直接变符号,否则偶数个相同的数会使得结果负负得正
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



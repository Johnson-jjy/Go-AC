package two_pointers

// 解法一: 排序,排序完后前后指针往中间扫,直到找到不对者
// 解法二: 单调栈 -> 待补充
// 解法三: 双指针 -> 本质上是要找到逆序的地方,从前往后和从后往前,然后判断中间段的min和max是否影响前后两端,即再扫一次
// 最后有三种解法的补充,其中双指针解法的最佳解是每次遇到逆序对时直接更新min和max,事实上这样一定能保证中间段的min和max都有记载,然后从首尾向中间扫即可
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	start := -1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {
			start = i
			break
		}
	}
	if start == -1 {
		return 0
	}
	stop := len(nums)
	for i := len(nums) - 1; i > start; i-- {
		if nums[i] < nums[i - 1] {
			stop = i
			break
		}
	}

	findMin := nums[start + 1]
	findMax := nums[start]
	for i := start; i <= stop; i++ {
		if nums[i] < findMin {
			findMin = nums[i]
		}
		if nums[i] > findMax {
			findMax = nums[i]
		}
	}
	for i := 0; i < start; i++ {
		if nums[i] > findMin {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i > stop; i-- {
		if nums[i] < findMax {
			stop = i
			break
		}
	}

	return stop - start + 1
}

// 排序 O(nlogn)
//func findUnsortedSubarray(nums []int) int {
//	sortedNums := make([]int, len(nums))
//	copy(sortedNums, nums)
//	sort.Slice(sortedNums, func(i, j int) bool {
//		return sortedNums[i] < sortedNums[j]
//	})
//	i, j := 0, len(nums)-1
//
//	// 1,2, 3,4
//	// 1 3  2 4
//	for i <= j && nums[i] == sortedNums[i] {
//		i++
//	}
//	for i <= j && nums[j] == sortedNums[j] {
//		j--
//	}
//	return j - i + 1
//}

// 单调栈 O（n）, O(n)
//func findUnsortedSubarray(nums []int) int {
//	stack := []int{0}
//	for i := 1; i < len(nums); i++ {
//		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//	l := 0
//	for l < len(stack) && stack[l] == l {
//		l++
//	}
//	stack = []int{len(nums) - 1}
//	for i := len(nums) - 2; i >= 0; i-- {
//		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//
//	r := len(nums) - 1
//
//	for len(nums)-r-1 < len(stack) && stack[len(nums)-r-1] == r {
//		r--
//	}
//
//	if l >= r {
//		return 0
//	}
//	return r - l + 1
//}

// O(n), O(1)
//func findUnsortedSubarray(nums []int) int {
//	curMin := math.MaxInt32
//	curMax := math.MinInt32
//	m := len(nums)
//	for i := 1; i < m; i++ {
//		if nums[i] < nums[i-1] {
//			if nums[i] < curMin {
//				curMin = nums[i]
//			}
//			if nums[i-1] > curMax {
//				curMax = nums[i-1]
//			}
//		}
//	}
//	if curMax == math.MinInt32 && curMin == math.MaxInt32 {
//		return 0
//	}
//	l, r := 0, m-1
//	for ; l < m && nums[l] <= curMin; l++ {
//	}
//	for ; r >= 0 && nums[r] >= curMax; r-- {
//	}
//
//	return r - l + 1
//}



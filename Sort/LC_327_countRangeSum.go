package Sort

// 思路: 计算个数 -> 归并 -> mSmS + calculate(根据规则) + m
// 重点理解: 对preSum数组的计数和对原数组的计数本质相同
// 待补充: 各种更复杂的数据结构
var count327, low, up int

func countRangeSum(nums []int, lower int, upper int) int {
	count327, low, up = 0, lower, upper
	pre := make([]int, len(nums) + 1)
	for i := 1; i <= len(nums); i++ {
		pre[i] = pre[i - 1] + nums[i - 1]
	}

	mergeSort(pre)

	return count327
}

func mergeSort327(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	p := len(nums) / 2
	left := mergeSort327(nums[:p])
	right := mergeSort327(nums[p:])

	l, r := 0, 0 // 注意: 考虑到left和right都是递增的,此处可以将此二者定义在函数外
	for _, v := range left {
		//fmt.Printf("%d\t%v\n", v, nums)
		for l < len(right) && right[l] - v < low {
			l++
		}
		for r < len(right) && right[r] - v <= up {
			r++
		}
		count327 += r - l
		//fmt.Printf("%d\t%d\t%d\t%v\n", l, r, count, nums)
	}

	return merge(left, right)
}

//func merge(left []int, right []int) []int {
//	i, j := 0, 0
//	m, n := len(left), len(right)
//	index := 0
//	result := make([]int, m + n)
//
//	for i < m && j < n {
//		if left[i] < right[j] {
//			result[index] = left[i]
//			index++
//			i++
//		} else {
//			result[index] = right[j]
//			index++
//			j++
//		}
//	}
//	for i < m {
//		result[index] = left[i]
//		index++
//		i++
//	}
//	for j < n {
//		result[index] = right[j]
//		index++
//		j++
//	}
//
//	return result
//}

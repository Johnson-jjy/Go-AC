package Sort

// 思路: 依然是归并排序经典题目 -> 注: 这一次不在归并的过程中排序了(逻辑过于复杂)
// 待补充: 树状数组
// 另: 末尾有一个初始时希望在归并过程排序的错误解法
var count int

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	count = 0
	mergeSort(nums)

	return count
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	p := len(nums)/2
	left := mergeSort(nums[:p])
	right := mergeSort(nums[p:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	result := make([]int, m + n)
	index := 0

	// 本质上, 类似于逆序对的题目都可以用形如 mergeSort,mergeSort,calculate,merge的方式解决
	// 只是若能直接在merge中解决,可节约calculate的复杂度
	for l,r := 0, 0; l < m; l++{
		for r < n && left[l] > 2 * right[r]{
			r++
		}
		count += r
	}

	for i < m && j < n {
		if left[i] <= right[j] {
			result[index] = left[i]
			index++
			i++
		} else {
			result[index] = right[j]
			index++
			j++
		}
	}

	for i < m {
		result[index] = left[i]
		index++
		i++
	}
	for j < n {
		result[index] = right[j]
		index++
		j++
	}
	return result
}


// 错误分析: 总是希望在归并过程中完成排序,但对于出现负数的情况不好处理
//var count int
//
//func reversePairs(nums []int) int {
//	if len(nums) <= 1 {
//		return 0
//	}
//	count = 0
//	mergeSort(nums)
//
//	return count
//}
//
//func mergeSort(nums []int) []int {
//	if len(nums) <= 1 {
//		return nums
//	}
//
//	p := len(nums)/2
//	left := mergeSort(nums[:p])
//	right := mergeSort(nums[p:])
//
//	return merge(left, right)
//}
//
//func merge(left []int, right []int) []int {
//	i, j := 0, 0
//	m, n := len(left), len(right)
//	result := make([]int, m + n)
//	index := 0
//	cur := 0
//
//	for i < m && j < n {
//		if left[i] < right[j] {
//			result[index] = left[i]
//			index++
//			i++
//		} else if left[i] == right[j] {
//			if right[j] < 0 {
//				count += m - i
//				result[index] = right[j]
//				index++
//				j++
//			} else {
//				result[index] = left[i]
//				index++
//				i++
//			}
//		} else {
//			if cur != m {
//				cur = i + findDouble(right[j], left[i:])
//				count += m - cur
//				//fmt.Printf("%d\t%v\t%v\t%d\t%d\n", i, left, right, cur, count)
//			}
//			result[index] = right[j]
//			index++
//			j++
//		}
//	}
//
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
//	return result
//}
//
//func findDouble(small int, arr []int) int {
//	target := small * 2
//	left := 0
//	right := len(arr) - 1
//	ans := len(arr)
//	for left <= right {
//		mid := left + (right - left)/2
//		if target < arr[mid] {
//			ans = mid
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return ans
//}

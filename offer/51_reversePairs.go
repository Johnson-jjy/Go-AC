package offer


// 解法一: 归并排序: 在merge中对逆序对计数, 注意计数的逻辑, 勿重复, 勿遗漏
// 解法二: 树状数组->待补充
var count int // 注意,LC_Oj中此值每次需设置为0

func reversePairs1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	count = 0
	nums = mergeSort(nums)

	return count
}

func mergeSort(nums []int) []int{
	if len(nums) <= 1 {
		return nums
	}// 注意此处的 <= 1, 若为 == 0的话,则会出现runtime error->left和right的分配会出问题

	p := len(nums) / 2
	left := mergeSort(nums[:p])
	right := mergeSort(nums[p:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	result := make([]int, m + n)
	index := 0

	for i < m && j < n {
		if left[i] <= right[j] {
			result[index] = left[i]
			index++
			i++
		} else {
			result[index] = right[j]
			index++
			j++
			count += (m - i)
		}
	}
	for i < m {
		result[index] = left[i]
		index++
		i++
		//count++
	}
	for j < n {
		result[index] = right[j]
		index++
		j++
	}

	return result
}



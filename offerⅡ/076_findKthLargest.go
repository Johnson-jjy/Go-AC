package offer_

// 数组中的第k大的数字
// 解一: 基于快排
// 1. 快排和归排, 在进递归时, 都需要对left<right的情况进行处理
// 2. 根据题目要求 (求第k大或者第k小) 合理修改target的值
// 3. 根据target和返回的pivot(right)值, 进行判断进而再递归
// 4. 在partition里, left <= right (原因: left和right都有可能取到; 类似二分)
// 5. nums[left] <= pivot; nums[right] > pivot; 交换start和right; 勿改此模板了, 针对target做文章足矣!
var target int

func findKthLargest(nums []int, k int) int {
	target = len(nums) - k
	quickSort(nums, 0, len(nums) - 1)

	//fmt.Println(nums)
	return nums[target]
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	p := partition(nums, start, end)
	//fmt.Println(nums, p)
	if p == target {
		return
	} else if p > target {
		quickSort(nums, start, p - 1)
	} else {
		quickSort(nums, p + 1, end)
	}
}

func partition(nums []int, start int, end int) int {
	pivot := nums[start]
	left := start + 1
	right := end
	for left <= right {
		for left <= right && nums[left] <= pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[start] = nums[start], nums[right]
	return right
}

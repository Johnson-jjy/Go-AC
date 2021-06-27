package main

import (
	"fmt"
)

// 左闭右开
func binary1(array []int, target int) int {
	var left, right, middle, count int
	
	left = 0
	right = len(array)
	count = 0

	fmt.Println("以下为左闭右开测试：")

	for left < right {
		count++
		middle = left + (right - left) / 2
		fmt.Printf("第%d次计算middle:%d;此时left为:%d;right为:%d\n",
			count, middle, left, right )

		if array[middle] > target {
			right = middle
		} else if array[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}


// 左闭右闭
func binary2(array []int, target int) int {
	var left, right, middle, count int

	left = 0
	right = len(array) - 1
	count = 0

	fmt.Println("以下为左闭右闭测试：")

	for left <= right {
		count++
		middle = left + (right - left) / 2
		fmt.Printf("第%d次计算middle:%d;此时left为:%d;right为:%d\n",
			count, middle, left, right )

		if array[middle] > target {
			right = middle - 1
		} else if array[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

// 1. 求等于x的最小的index，不存在返回-1
func binaryFindEqualMinIndex(array []int, target int) int {
	var left, middle, right, ans int

	left = 0
	right = len(array) - 1
	ans = -1

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] == target {
			ans = middle
			right = middle - 1
		} else if array[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return ans
}

// 2. 求等于x的最大的index，不存在返回 -1
func binaryFindEqualMaxIndex(array []int, target int) int {
	var left, right, middle, ans int

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] == target {
			ans = middle
			left = middle + 1
		} else if array[middle] > target {
			right = middle + 1
		} else {
			left = middle + 1
		}
	}

	return ans
}

// 3. 求小于x的最大的index
func binaryFindLessMaxIndex(array []int, target int) int {
	var left, right, middle int

	left = 0
	right = len(array) - 1

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] >= target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return right
}

// 4. 求大于x的最小的index
func binaryFindMoreMinIndex(array []int, target int) int {
	var left, right, middle int

	left = 0
	right = len(array) - 1

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}

// 5. 求大于等于x的最小的index
func binaryFindMoreEqualMinIndex(array []int, target int) int {
	var left, right, middle int

	left = 0
	right = len(array) - 1

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] >= target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return left
}

// 6. 求小于等于x的最大的index
func binaryFindLessEqualMinIndex(array []int, target int) int {
	var left, right, middle int

	left = 0
	right = len(array) - 1

	for left <= right {
		middle = left + (right - left) / 2
		if array[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return right
}

// 左开右开
func binary3(array []int, target int) int {
	var left, right, middle, count int

	left = -1
	right = len(array)
	count = 0

	fmt.Println("以下为左开右开测试：")

	for left + 1 != right {
		count++
		middle = left + (right - left) / 2

		fmt.Printf("第%d次计算middle:%d;此时left为:%d;right为:%d\n",
			count, middle, left, right )

		if array[middle] > target {
			right = middle
		} else if array[middle] < target {
			left = middle
		} else {
			return middle
		}
	}

	return -1
}

func main()  {
	test_odd := []int{0, 1, 2, 3, 4}
	test_even := []int{0, 1, 2, 3, 4, 5}

	binary1(test_odd, 4)
	binary1(test_even, 4)

	binary2(test_odd, 4)
	binary2(test_even, 4)

	binary3(test_odd, 4)
	binary3(test_even, 4)
}
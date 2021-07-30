package offer

// 思路1: 快排,但不用严格排序
// 思路2: 堆排,待补充
var target int

func getLeastNumbers(arr []int, k int) []int {
	target = k - 1
	quickSort(arr, 0, len(arr) - 1)

	return arr[ : k]
}

func quickSort(arr []int, left int, right int) {
	if left < right {
		p := partition(arr, left, right)
		if p == target {
			return
		} else if p > target {
			quickSort(arr, left, p - 1)
		} else {
			quickSort(arr, p + 1, right)
		}
	}
}

func partition(arr[] int, left int, right int) int {
	pivot := arr[left]
	i := left + 1
	j := right
	for i <= j {
		for i <= j && arr[i] <= pivot {
			i++
		}
		for i <= j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j], arr[left] = arr[left], arr[j]

	return j
}

// 注: partition这样写是不行的
// 这样写无法满足前K个一定恰好是前K个最小的数
//func quickSort(arr []int, left int, right int) {
//	if left < right {
//		p := partition(arr, left, right)
//		if p >= target {
//			return
//		} else {
//			quickSort(arr, p + 1, right)
//		}
//	}
//}
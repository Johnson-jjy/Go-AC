package cig

// 经典快排
// 注意 i, j 以及对p的选择
var target int

func smallestK(arr []int, k int) []int {
	target = k - 1
	quickSort(arr, 0, len(arr) - 1)

	return arr[:k]
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

func partition(arr []int, left int, right int) int {
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

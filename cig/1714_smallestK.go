package cig

// 经典快排
// 注意 i, j 以及对p的选择
// 1.quickSort不需要返回数组,go对切片的修改会影响到底层; 2.进quickSort时需要判断start和end的大小关系;
// 3.partition不能传入子数组,否则无法判断得到的p值与target的关系; 4.注意,一定要在quickSort之前对target进行赋值,否则每次目标target都是0
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

//var target int
//
//func smallestK(arr []int, k int) []int {
//	quicksort(arr, 0, len(arr) - 1)
//	target = k - 1 -> !target必须在quicksort之前,不然就是零值了!
//	return arr[:k]
//}
//
//func quicksort(arr []int, start int, end int) {
//	if start >= end {
//		return
//	}
//	p := partition(arr, start, end)
//	fmt.Println(arr, p)
//	if p == target {
//		return
//	} else if p > target {
//		quicksort(arr,start, p - 1)
//	} else {
//		quicksort(arr, p + 1, end)
//	}
//}
//
//func partition(arr []int, start int, end int) int {
//	left := start + 1
//	right := end
//	flag := arr[start]
//	//fmt.Println(left, right, flag)
//	for left <= right {
//		for left <= right && arr[left] <= flag {
//			left++
//		}
//		for left <= right && arr[right] > flag {
//			right--
//		}
//		if left < right {
//			arr[left], arr[right] = arr[right], arr[left]
//		}
//	}
//	arr[right], arr[start] = arr[start], arr[right]
//	// fmt.Println(arr, right)
//	return right
//}
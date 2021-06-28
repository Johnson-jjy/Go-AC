package Binary_Search

func peakIndexInMountainArray(arr []int) int {
	var left, right int
	left = 0
	right = len(arr) - 1
	mid := 0
	for left <= right {
		mid = (right - left) / 2 + left
		if arr[mid - 1] < arr[mid] && arr[mid] > arr[mid + 1] {
			return mid
		} else if arr[mid - 1] > arr[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return mid
}

// 官方题解：更快，涉及sort源码，需理解
//func peakIndexInMountainArray(arr []int) int {
//	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
//}

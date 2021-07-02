package Binary_Search

func peakIndexInMountainArray(arr []int) int {
	left := 1
	right := len(arr) - 2

	ans := right
	for left <= right {
		mid := left + (right - left) / 2

		if arr[mid] < arr[mid + 1] {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}

	return ans
}

// 官方题解：更快，涉及sort源码，需理解
//func peakIndexInMountainArray(arr []int) int {
//	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
//}

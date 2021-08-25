package Binary_Search

import (
	"math"
	"sort"
)

// 找到K个最接近的元素
// 解一: 看到排序数组 -> 想二分, 先通过二分找到第一个比x的数的坐标; 2.然后想左右延展即可
func findClosestElements1(arr []int, k int, x int) []int {
	fmax := len(arr)
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right - left)/2
		//fmt.Println(mid, arr[mid], fmax)
		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
			fmax = mid
		}
	}
	if fmax == len(arr) {
		return arr[len(arr) - k : ]
	}
	if fmax == -1 || fmax == 0 {
		return arr[:k]
	}
	fmin := fmax - 1
	index := 0
	res := make([]int, k)
	for index < k {
		maxAbs := math.MaxInt32
		minAbs := math.MaxInt32
		if fmax < len(arr) {
			maxAbs = abs(x, arr[fmax])
		}
		if fmin >= 0 {
			minAbs = abs(x, arr[fmin])
		}
		if maxAbs < minAbs {
			res[index] = arr[fmax]
			fmax++
		} else {
			res[index] = arr[fmin]
			fmin--
		}
		index++
	}

	sort.Ints(res)
	return res
}

// 解二: 思想待深入理解
func findClosestElements2(arr []int, k int, x int) []int {
	start := 0
	end   := len(arr) - k - 1

	for start <= end {
		mid := start + (end - start) / 2

		if x - arr[mid] > arr[mid+k] - x{
			start = mid + 1
		}else{
			end = mid - 1
		}
	}

	return arr[start:start+k]
}

//func abs(a, b int) int {
//	if a > b {
//		return a - b
//	}
//	return b - a
//}

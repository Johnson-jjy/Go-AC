package main

import (
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 给定一个整数数组arr和一个数K，返回数组中小于K的最大的两个数之和。如果不存在小于K的两数之和，返回-1。
 * @param arr int整型一维数组 数组
 * @param k int整型 相加的上限
 * @return int整型
 */
func getMax( arr []int ,  k int ) int {
	// write code here
	sort.Ints(arr)
	if arr[0] >= k {
		return -1
	}
	ans := -1
	left := 0
	right := len(arr) - 1
	for left < right {
		cur := arr[left] + arr[right]
		if cur < k {
			if ans < cur {
				ans = cur
			}
			left++
		} else {
			right--
		}
	}

	return ans
}
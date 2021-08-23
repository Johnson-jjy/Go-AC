//package main
//
//import "fmt"
//
///**
//* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
//*
//*
//* @param heights int整型一维数组
//* @return int整型一维数组
//*/
//
//// 逛街
//func findBuilding( heights []int ) []int {
//	// write code here
//	n := len(heights)
//	res := make([]int, n)
//	if n == 0 {
//		return res
//	}
//	if n == 1 {
//		res[0] = 1
//		return res
//	}
//	left := make([]int, n)
//	right := make([]int, n)
//
//	for i := n - 2; i >= 0; i-- {
//		low :=  heights[i + 1]
//		cur := 1
//		for j := i + 1; j < n; j++ {
//			if heights[j] > low {
//				cur++
//				low = heights[j]
//			}
//		}
//		right[i] = cur
//	}
//	fmt.Println(right)
//
//	for i := 1; i < n; i++ {
//		low :=  heights[i - 1]
//		cur := 1
//		for j := i - 1; j >= 0; j-- {
//			if heights[j] > low {
//				cur++
//				low = heights[j]
//			}
//		}
//		left[i] = cur
//	}
//	fmt.Println(left)
//
//	for i := 0; i < n; i++ {
//		res[i] = left[i] + right[i] + 1
//	}
//
//	return res
//}

func main() {
	heights := []int{5,3,8,3,2,5}
	fmt.Println(findBuilding(heights))
}


package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param heights int整型一维数组
 * @return int整型一维数组
 */

// 逛街
func findBuilding( heights []int ) []int {
	// write code here
	n := len(heights)
	res := make([]int, n)
	if n == 0 {
		return res
	}
	if n == 1 {
		res[0] = 1
		return res
	}
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		right[i] = len(stack)
		for len(stack) > 0 && stack[len(stack) - 1] < heights[i] {
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, heights[i])
	}
	fmt.Println(right)

	stack = make([]int, 0)
	for i := 0; i < n; i++ {
		left[i] = len(stack)
		for len(stack) > 0 && stack[len(stack) - 1] < heights[i]{
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, heights[i])
	}
	fmt.Println(left)

	for i := 0; i < n; i++ {
		res[i] = left[i] + right[i] + 1
	}

	return res
}

func main() {
	heights := []int{5,3,8,3,2,5}
	fmt.Println(findBuilding(heights))
}
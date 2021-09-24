package main

import (
	"fmt"
	"math"
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution3(A []int) int {
	// write your code in Go 1.4
	n := len(A)
	nums := make([]int, n)
	copy(nums, A)
	sort.Ints(nums)

	res := 0
	max := math.MinInt32
	for i := 0; i < n; i++ {
		left := 0
		right := n - 1
		index := n
		for left <= right {
			mid := left + (right - left)/2
			if nums[mid] == A[i] {
				index = mid
				break
			} else if nums[mid] > A[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if index > max {
			max = index
		}
		if max == i {
			res++
			max = math.MinInt32
		}
	}

	return res
}

func main()  {
	fmt.Println(Solution3([]int{4, 3, 2, 6, 1}))
}
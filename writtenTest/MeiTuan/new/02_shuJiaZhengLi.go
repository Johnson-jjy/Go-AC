package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	res := 1
	for i := n - 1; i >= 0; i-- {
		target := a[i]
		left := 0
		right := n - 1
		ans := n
		for left <= right {
			mid := left + (right - left)/2
			if b[mid] >= target {
				ans = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		cur := (n - ans) - (n - i) + 1
		//fmt.Println(res, ans, b[ans], a[i], cur)
		if cur < 0 {
			fmt.Println(0)
			return
		}
		res *= cur
	}

	fmt.Println(res)
}

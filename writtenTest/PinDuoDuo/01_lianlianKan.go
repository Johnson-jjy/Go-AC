package main

import (
	"fmt"
	"sort"
)

func main()  {
	var T, n, m, k int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&n, &m, &k)
		res := 0
		red := make([]int, n)
		blue := make([]int, m)
		for j := 0; j < n; j++ {
			fmt.Scan(&red[j])
		}
		for j := 0; j < m; j++ {
			fmt.Scan(&blue[j])
		}
		sort.Ints(red)
		sort.Ints(blue)
		left := 0
		right := 0
		for left < n && right < m {
			if red[left] > blue[right] {
				if red[left] - blue[right] <= k {
					res++
					left++
					right++
				} else {
					right++
				}
			} else {
				if blue[right] - red[left] <= k {
					res++
					left++
					right++
				} else {
					left++
				}
			}
		}
		fmt.Println(res)
	}
}

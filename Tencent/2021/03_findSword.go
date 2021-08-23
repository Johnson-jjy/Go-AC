package main

import (
	"fmt"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var n, w int
		fmt.Scan(&n)
		fmt.Scan(&w)

		store := make([]int, n)
		for j := 0;j < n; j++ {
			fmt.Scan(&store[j])
		}
		sort.Ints(store)

		res := 0
		for k := n - 1; k >= 0; k-- {
			if store[k] > w {
				continue
			}
			res++
			need := w - store[k]
			if need > 0 {
				// 二分怎么处理原地hash的情况?先记录一下,有时间再来补
				for m := k - 1; m >= 0; m-- {
					if store[m] <= need && (store[m] + store[k]) % 2 == 0 {
						store[m] += w
						break
					}
				}
			}
		}

		fmt.Println(res)
	}
}

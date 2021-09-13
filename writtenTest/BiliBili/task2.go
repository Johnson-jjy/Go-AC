package main

import (
	"fmt"
	"math"
)

func main()  {
	var m, n int
	fmt.Scan(&m, &n)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	pointer := make([]int, m)
	cur := math.MaxInt32
	index := -1
	for  {
		for i := 0; i < m; i++ {
			if pointer[i] < n {
				if arr[i][pointer[i]] < cur {
					cur = arr[i][pointer[i]]
					index = i
				}
			}
		}
		if index != -1 {
			fmt.Println(cur)
			pointer[index]++
			cur = math.MaxInt32
			index = -1
		} else {
			break
		}
	}

}
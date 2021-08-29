package main

import (
	"fmt"
)


func main()  {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	store := make([]int, 31)
	check := make([]bool, 31)
	for i := 0; i < n; i++ {
		if !check[arr[i]] {
			check[arr[i]] = true
			for j := arr[i] + 1; j <= 30; j++ {
				store[j]++
			}
		}
		res += store[arr[i]]
	}

	fmt.Println(res)
}



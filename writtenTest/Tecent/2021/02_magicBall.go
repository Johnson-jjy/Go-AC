package main

import (
	"fmt"
	"sort"
)

var max int

func main() {
	var T int
	fmt.Scan(&T)
	max = 1000000007
	store := make([]int, 1001)
	store[0] = 1
	for i := 1; i <= 1000; i++ {
		store[i] = store[i - 1] * 2 % max
	}

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}
		sort.Ints(arr)
		res := 0
		for k := len(arr) - 1; k >= 0; k-- {
			cur := arr[k]
			times := store[k]
			res = (res + cur * times) % max
		}
		fmt.Println(res)
	}
}
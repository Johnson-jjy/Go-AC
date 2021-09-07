package main

import "fmt"

func main()  {
	var n, k int
	fmt.Scan(&n, &k)
	store := make([][]int, n)
	res := make([][]int, n * k)
	for i := 0; i < n; i++ {
		store[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&store[i][j])
		}
	}
	for i := 0; i < n * k; i++ {
		res[i] = make([]int, n * k)
		for j := 0; j < n * k; j++ {
			res[i][j] = store[i/k][j/k]
			fmt.Print(res[i][j])
			if j == n * k - 1 {
				fmt.Printf("\n")
			} else {
				fmt.Print(" ")
			}
		}
		//fmt.Println(res[i])
	}

	//fmt.Println(res)
}

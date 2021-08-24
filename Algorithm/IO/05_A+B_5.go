package main

import "fmt"

/*
	2
	4 1 2 3 4
	5 1 2 3 4 5
*/

func main() {
	var n, m, cur int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		sum := 0
		for j := 0; j < m; j++ {
			fmt.Scan(&cur)
			sum += cur
		}
		fmt.Println(sum)
	}
}
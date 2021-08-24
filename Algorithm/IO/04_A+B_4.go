package main

import "fmt"

/*
	4 1 2 3 4
	5 1 2 3 4 5
	0
*/

func main() {
	var n, cur int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&cur)
			sum += cur
		}
		fmt.Println(sum)
	}
}

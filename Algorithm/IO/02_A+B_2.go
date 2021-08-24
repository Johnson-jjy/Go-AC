package main

import "fmt"

/*
	2
	1 5
	10 20
*/

func main() {
	var a, b, n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
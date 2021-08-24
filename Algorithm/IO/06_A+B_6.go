package main

import (
	"fmt"
	"io"
)

/*
	4 1 2 3 4
	5 1 2 3 4 5
*/

func main() {
	var n, cur int
	for {
		if _, err := fmt.Scan(&n); err != io.EOF {
			sum := 0
			for i := 0; i < n; i++ {
				fmt.Scan(&cur)
				sum += cur
			}
			fmt.Println(sum)
		} else {
			break
		}
	}
}

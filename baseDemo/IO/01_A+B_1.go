package main

import (
	"fmt"
	"io"
)

/*
	1	5
	10	20
*/

func main() {
	var a, b int
	for {
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println(a + b)
		} else {
			break
		}
	}
}

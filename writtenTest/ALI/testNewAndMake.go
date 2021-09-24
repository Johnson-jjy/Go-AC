package main

import "fmt"

func main()  {
	a := new([]int)
	b := make([]int, 0)

	fmt.Println(*a, b)
}

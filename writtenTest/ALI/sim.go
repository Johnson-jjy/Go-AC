package main

import "fmt"

func main()  {
	var n, a, b int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		res := 2
		for res < i {
			res *= 2
		}
		if res * 2 < b {
			res = 0
		}
		fmt.Println(res)
	}
}

//func count(n int) int {
//	c := 0
//	for n != 0 {
//		c++
//		n = n & (n - 1)
//	}
//
//	return c
//}

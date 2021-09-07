package main

import "fmt"

func main()  {
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		res := 0
		for j := 1; j * j < n; j++ {
			if n % j == 0 {
				if gcd(j, n/j) == 1 {
					res++
				}
			}
		}
		fmt.Println(res)
	}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

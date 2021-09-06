package main

import "fmt"

func main()  {
	var n, l, r int
	fmt.Scan(&n, &l, &r)

	arr := make([]int, 0)
	var dfs func(n int)
	dfs = func(n int) {
		if n < 2 {
			arr = append(arr, n)
			//fmt.Println(arr)
			return
		}
		f := n/2
		s := n % 2
		t := n/2
		dfs(f)
		dfs(s)
		dfs(t)
	}

	dfs(n)
	res := 0
	//fmt.Println(arr)
	for i := l - 1; i <= r - 1; i++ {
		res += arr[i]
	}
	fmt.Println(res)
}

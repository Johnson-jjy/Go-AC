package main

import "fmt"

func main()  {
	var n, q int
	fmt.Scan(&n, &q)
	open := make([][]bool, n)
	shut := make([][]bool, n)
	for i := 0; i < n; i++ {
		open[i] = make([]bool, n)
		shut[i] = make([]bool, n)
	}

	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		var t int
		for j := 0; j < num; j++ {
			fmt.Scan(&t)
			//fmt.Println(n, num, t)
			open[i][t - 1] = true
			shut[t - 1][i] = true
		}
	}

	//fmt.Println(open)
	//fmt.Println(shut)

	state := make([]bool, n)
	var choice, machine int

	var openDFS func(cur int)
	openDFS = func(cur int) {
		state[cur] = true
		for j := 0; j < n; j++ {
			if open[cur][j] && !state[j] {
				openDFS(j)
			}
		}
	}

	var shutDFS func(cur int)
	shutDFS = func(cur int) {
		state[cur] = false
		for j := 0; j < n; j++ {
			if shut[cur][j] && state[j] {
				shutDFS(j)
			}
		}
	}

	count := func() int {
		res := 0
		for i := 0; i < n; i++ {
			if state[i] {
				res++
			}
		}
		return res
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&choice)
		fmt.Scan(&machine)
		if choice == 1 {
			openDFS(machine - 1)
		} else {
			shutDFS(machine - 1)
		}
		fmt.Println(count())
	}
}


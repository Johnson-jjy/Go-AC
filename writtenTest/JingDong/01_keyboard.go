package main

import "fmt"

func main()  {
	var n, m, x, y, z int
	fmt.Scan(&n, &m, &x, &y, &z)
	var str string
	grid := make(map[byte][]int, 0)
	var cur string
	for i := 0; i < n; i++ {
		fmt.Scan(&cur)
		//fmt.Println("?", cur)
		for j := 0; j < m; j++ {
			grid[cur[j]] = []int{i, j}
		}
	}
	fmt.Scan(&str)

	res := 0

	start := []int{0, 0}
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] == str[i - 1] {
			res += z
			continue
		}
		startX := start[0]
		startY := start[1]
		curX := grid[str[i]][0]
		curY := grid[str[i]][1]
		moveX := abs(curX, startX) * x
		moveY := abs(curY, startY) * x
		if moveX != 0 && moveY != 0 {
			res += y
		}
		res += moveX + moveY + z
		start = []int{curX, curY}
	}

	fmt.Println(res)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

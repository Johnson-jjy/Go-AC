package island

// 统计封闭岛屿的数目->读懂题意很关键!
// 关键点: 要预先对四边的情况进行预处理->预处理也是走递归
func closedIsland(grid [][]int) int {
	res := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])

	var backtrack func(i int, j int)
	backtrack = func(i int, j int) {
		grid[i][j] = 1

		for k := 0; k < 4; k++ {
			nextI := i + directions[k][0]
			nextJ := j + directions[k][1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == 0 {
				backtrack(nextI, nextJ)
			}
		}
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			backtrack(0, j)
		}
		if grid[m - 1][j] == 0 {
			backtrack(m - 1, j)
		}
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			backtrack(i, 0)
		}
		if grid[i][n - 1] == 0 {
			backtrack(i, n - 1)
		}
	}

	for i := 1; i < m - 1; i++ {
		for j := 1; j < n - 1; j++ {
			if grid[i][j] == 0 {
				res++
				backtrack(i, j)
			}
		}
	}

	return res
}

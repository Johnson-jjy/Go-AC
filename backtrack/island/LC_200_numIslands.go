package island

// 岛屿数量: 岛屿类题目母题
// 1.回溯->每次用directions; 2.修改grid或visited; 3.针对源点操作,计数或者添加修改等
// 注意: 关于nextI和nextJ的地方不要用错了
func numIslands(grid [][]byte) int {
	res := 0
	// 注意directions为二维数组
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])

	var backtrack func(i int, j int)
	backtrack = func(i int, j int) {
		grid[i][j] = '0'

		for k := 0; k < 4; k++ {
			nextI := i + directions[k][0]
			nextJ := j + directions[k][1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == '1' {
				backtrack(nextI, nextJ)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				backtrack(i, j)
			}
		}
	}

	return res
}

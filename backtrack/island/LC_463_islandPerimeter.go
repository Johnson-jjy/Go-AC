package island

// 岛屿的周长: 抓只有一个岛屿这一关键信息-> 利用数学逻辑求解即可
func islandPerimeter(grid [][]int) int {
	num := 0
	near := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				num++
				if j > 0 && grid[i][j - 1] == 1 {
					near++
				}
				if i > 0 && grid[i - 1][j] == 1 {
					near++
				}
			}
		}
	}

	return num * 4 - near * 2
}

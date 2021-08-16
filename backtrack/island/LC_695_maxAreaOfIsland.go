package island

// 岛屿最大面积
// 注意每次求得的sum,因为会返回,故直接赋给if中的sum即可,不需要+=
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m := len(grid)
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	if m == 0 {
		return res
	}
	n := len(grid[0])

	var backtrack func(sum int, i int, j int) int
	backtrack = func(sum int, i int, j int) int {
		grid[i][j] = 0
		sum++

		for k := 0; k < 4; k++ {
			nextI := i + directions[k][0]
			nextJ := j + directions[k][1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == 1 {
				sum = backtrack(sum, nextI, nextJ)
			}
		}

		return sum
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur := backtrack(0, i, j)
				if cur > res {
					res = cur
				}
			}
		}
	}

	return res
}

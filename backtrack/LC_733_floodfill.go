package backtrack

// 图像渲染
// DFS标准洪泛
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	if image[sr][sc] == newColor {
		return image
	}
	m := len(image)
	n := len(image[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	oldColor := image[sr][sc]

	var dfs func(curX, curY int)
	dfs = func(curX, curY int) {
		visited[curX][curY] = true
		image[curX][curY] = newColor

		for i := 0; i < 4; i++ {
			nX := curX + directions[i][0]
			nY := curY + directions[i][1]
			if nX >= 0 && nX < m && nY >= 0 && nY < n && !visited[nX][nY] && image[nX][nY] == oldColor {
				dfs(nX, nY)
			}
		}
	}

	dfs(sr, sc)

	return image
}

package offer

// 机器人的运动范围

// 思路: DFS(也可BFS,可达点用queue保存)
// 注: 1.有些点满足k但不可达; 2.题目的判点条件; 3.visited数组勿忘
func movingCount(m int, n int, k int) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	count := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	sum := func(x, y int) int {
		res := 0
		for x > 0 {
			res += x % 10
			x /= 10
		}
		for y > 0 {
			res += y % 10
			y /= 10
		}
		return res
	}

	var dfs func(curX int, curY int)
	dfs = func(curX int, curY int) {
		visited[curX][curY] = true
		count++

		for i := 0; i < 4; i++ {
			nX := curX + directions[i][0]
			nY := curY + directions[i][1]
			if nX >= 0 && nX < m && nY >= 0 && nY < n && sum(nX, nY) <= k && !visited[nX][nY] {
				dfs(nX, nY)
			}
		}
	}

	dfs(0, 0)

	return count
}

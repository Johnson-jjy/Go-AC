package path

// 不同路径Ⅱ
// 注:不要忘记初始化
func uniquePathsWithObstacles63(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1{
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[m - 1][n - 1] == 1 {
		return 0
	}

	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	dp[0][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i - 1][j - 1] == 1 {
				continue
			}
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	//fmt.Println(obstacleGrid, dp)
	return dp[m][n]
}

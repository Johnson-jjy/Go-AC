package greedy

// 逃脱阻碍者: 本质贪心; 都按往终点走的方式,计数即可; 注意,所有ghosts都是同时出发,而不存在延迟i步后再走的情况
// 求距离: 曼哈顿距离
func escapeGhosts(ghosts [][]int, target []int) bool {
	m := len(ghosts)

	start := []int{0, 0}
	need := abs(start[0], target[0]) + abs(start[1], target[1])

	for i := 0; i < m; i++ {
		cur := abs(ghosts[i][0], target[0]) + abs(ghosts[i][1], target[1])
		if cur <= need {
			return false
		}
	}

	return true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

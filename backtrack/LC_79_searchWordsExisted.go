package backtrack

// 单词搜索: 1. 每个点都可以是源点,两次for -> 类似岛屿类题目;
// 2. 注意,实际上index不会超过word的长度,因为到最后一个字符时,判断正确则返回true
// 3. 无论是对每个源点的判断还是对for的判断,都要if backtrack -> 一旦true则返回true
// 4. 不要忘了在每次的最后更新visited的状态
func exist(board [][]byte, word string) bool {
	var directions = [][]int {{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var backtrack func(i int, j int, index int) bool
	backtrack = func(i int, j int, index int) bool {
		//fmt.Println(i, j, index)
		if board[i][j] != word[index] {
			return false
		}

		if index == len(word) - 1 {
			return true
		}

		visited[i][j] = true

		for k := 0; k < 4; k++ {
			nextI := i + directions[k][0]
			nextJ := j + directions[k][1]
			if nextI >= 0 && nextI < m && nextJ >= 0 &&  nextJ < n && !visited[nextI][nextJ] {
				if backtrack(nextI, nextJ, index + 1) {
					return true
				}
			}
		}

		visited[i][j] = false

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

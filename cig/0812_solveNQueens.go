package cig

// 八皇后: 注意本题需要自己构建一个board
// 注意,for循环里不能声明多个变量
// 另: 待补充 -> 位运算优化版本
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var check func(i int, j int) bool
	check = func(i int, j int) bool {
		for curI := 0; curI < i; curI++ {
			if board[curI][j] == 'Q' {
				return false
			}
		}

		curI := i
		curJ := j
		for curI >= 0 && curJ >= 0 {
			if board[curI][curJ] == 'Q' {
				return false
			}
			curI--
			curJ--
		}

		curI, curJ = i, j
		for curI >= 0 && curJ < n{
			if board[curI][curJ] == 'Q' {
				return false
			}
			curI--
			curJ++
		}

		return true
	}

	var backtrack func(k int)
	backtrack = func(k int) {
		if k == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				cur := string(board[i])
				tmp[i] = cur
			}
			res = append(res, tmp)
			return
		}

		for j := 0; j < n; j++ {
			if check(k, j) {
				board[k][j] = 'Q'
				backtrack(k + 1)
				board[k][j] = '.'
			}
		}
	}

	backtrack(0)

	return res
}

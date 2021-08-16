package backtrack

// N皇后Ⅱ: 本题只需要返回组数,故用一个res计数即可,其余和51相同
func totalNQueens(n int) int {
	//directions := []int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := 0
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	check := func(i int, j int) bool {
		for cur := 0; cur < i; cur++ {
			if board[cur][j] == 'Q' {
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

		curI = i
		curJ = j
		for curI >= 0 && curJ < n {
			if board[curI][curJ] == 'Q' {
				return false
			}
			curI--
			curJ++
		}

		return true
	}

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			res++
			return
		}

		for j := 0; j < n; j++ {
			//nextI := i + directions[k]
			//nextJ := j + directions[k]
			if check(i, j){
				board[i][j] = 'Q'
				backtrack(i + 1)
				board[i][j] = '.'
			}
		}
	}

	backtrack(0)

	return res
}

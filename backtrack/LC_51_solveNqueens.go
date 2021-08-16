package backtrack

// N皇后1: 要返回相应的结果,故需要用二维strings数组存
// 注: N皇后的题目不需要和flood fill不同, 不需要用到directions存,只是每次对该层节点是否可用做一个判断
// 优化方向: 所有网格类判断格点是否满足某种规律的情况,都可以用位图进行优化
// 注: 判断是1.纵向; 2.往左上; 3.往右上
func solveNQueens(n int) [][]string {
	//directions := []int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := make([][]string, 0)
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
			tmp := make([]string, n)
			for k := 0; k < n; k++ {
				tmp[k] = string(board[k])
			}
			res = append(res, tmp)
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

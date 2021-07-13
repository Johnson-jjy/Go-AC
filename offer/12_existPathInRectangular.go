package offer

var visited [][]bool
var m int
var n int
var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// 标准的 回溯Flood Fill -> 二维数组定义 + 源点出发 + true值返回 + visited更新
// go语法声明二维数组, 勿错!
// 注意可从每个源点出发,切忌总从0,0出发
// 找到了true以后要提前返回
// 对于visited数组要及时更新
// 尽量传指针节约空间 -> 注:也可直接修改原数组(回溯的时候再重置即可) (之后有闭包+修改原数组的写法)
func exist1(board [][]byte, word string) bool {
	m = len(board)
	n = len(board[0])
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(&board, i, j, 0, &word) {
				return true
			}
		}
	}

	return false
}

func backtrack(board *[][]byte, curI int, curJ int, curIndex int, s *string) bool {
	if curI < 0 || curI >= m || curJ < 0 || curJ >= n {
		return false
	}
	if visited[curI][curJ] {
		return false
	}
	//fmt.Printf("before:%d\t%d\t%d\n", curI, curJ, curIndex)
	if curIndex >= len(*s) || (*board)[curI][curJ] != (*s)[curIndex] {
		return false
	}
	if curIndex == (len(*s) - 1) && (*board)[curI][curJ] == (*s)[curIndex] {
		return true
	}
	visited[curI][curJ] = true
	//fmt.Printf("after:%d\t%d\t%d\n", curI, curJ, curIndex)

	for i := 0; i < 4; i++ {
		nextI := curI + directions[i][0]
		nextJ := curJ + directions[i][1]
		if backtrack(board, nextI, nextJ, curIndex + 1, s) {
			return true
		}
	}
	visited[curI][curJ] = false
	return false
}

func exist2(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if !(i >= 0 && i < len(board)) ||
			!(j >= 0 && j < len(board[0])) ||
			board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		board[i][j] = ' '
		res := dfs(i+1, j, k+1) ||
			dfs(i-1, j, k+1) ||
			dfs(i, j+1, k+1) ||
			dfs(i, j-1, k+1)
		board[i][j] = word[k]
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
package backtrack

// 判断有效数独
// 解一: 一次遍历 + 哈希计数 -> 注:不需要用map存,直接用数组即可
func isValidSudoku1(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0 ; i < 9; i++ {
		for j := 0 ; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				index_box := (i/3) * 3 + j/3
				if row[i][num] == 1 {
					return false
				}else{
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				}else{
					col[j][num] = 1
				}
				if sbox[index_box][num] == 1 {
					return false
				}else {
					sbox[index_box][num] = 1
				}
			}
		}
	}
	return true
}

// 解二: 位图
func isValidSudoku2(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3  // 3x3的块索引号
			if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
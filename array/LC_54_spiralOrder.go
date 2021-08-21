package array

// 螺旋矩阵: -> 模拟
// 法一: 最基础的模拟: 1.每次到达边界后,行和列都要变化一位,否则会有重复数字; 2.每一轮要对circle进行计数
// 法二: 对left和right进行框定; 待补充
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	need := m * n
	index := 0
	res := make([]int, need)
	i := 0
	j := 0
	circle := 0

	for index < need {
		for index < need && j < n - circle {
			res[index] = matrix[i][j]
			index++
			j++
		}
		j--
		i++
		for index < need && i < m - circle {
			res[index] = matrix[i][j]
			index++
			i++
		}
		i--
		j--
		for index < need && j >= 0 + circle {
			res[index] = matrix[i][j]
			index++
			j--
		}
		j++
		i--
		for index < need && i >= 0 + circle + 1 {
			res[index] = matrix[i][j]
			index++
			i--
		}
		i++
		j++
		circle++
	}

	return res
}

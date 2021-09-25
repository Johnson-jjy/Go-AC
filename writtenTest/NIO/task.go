package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param N int整型 任务个数
 * @param relations int整型二维数组 依赖关系
 * @return int整型
 */
func calcEffort( N int ,  relations [][]int ) int {
	// write code here
	res := 0
	if len(relations) == 0 {
		return res
	}
	need := make([][]bool, N + 1)
	//let := make([][]bool, N + 1)
	for i := 0; i <= N; i++ {
		need[i] = make([]bool, N + 1)
		//let[i] = make([]bool, N + 1)
	}

	//fmt.Println(relations)

	m := len(relations)
	for i := 0; i < m; i++ {
		x := relations[i][0]
		y := relations[i][1]
		need[y][x] = true
		//let[x][y] = true
		if need[x][y] {
			return -1
		}
	}

	count := 0
	states := make([]bool, N + 1)
	for count != N {
		res++
		//fmt.Println(states)
		cur := make([]int, 0)

		for i := 1; i <= N; i++ {
			if states[i] {
				continue
			}
			//fmt.Println("?", res, i)
			can := true
			for j := 1; j <= N; j++ {
				if need[i][j] {
					//fmt.Println(i, j, need[i])
					can = false
					break
				}
			}
			if can {
				cur = append(cur, i)
			}
		}

		if len(cur) == 0 {
			//fmt.Println(need, states)
			return -1
		}
		for i := 0; i < len(cur); i++ {
			index := cur[i]
			states[index] = true
			for j := 1; j <= N; j++ {
				need[j][index] = false
			}
		}
		count += len(cur)
	}

	return res
}

func main()  {
	relations := [][]int{{1,3}, {2,3}}
	N := 3
	fmt.Println(calcEffort(N, relations))
}
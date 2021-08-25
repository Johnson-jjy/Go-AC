package main

import (
	"fmt"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回指定的点所在的多边形序号。
 * @param points int整型二维数组 2维平面的候选点坐标。
 * @param polygens int整型二维数组 多边形顶点坐标索引。
 * @param positions int整型二维数组 需要计算的指定点坐标。
 * @return int整型一维数组
 */
func PointInPolygen( points [][]int ,  polygens [][]int ,  positions [][]int ) []int {
	// write code here
	n := len(positions)
	res := make([]int, n)
	m := len(polygens)
	for i := 0; i < n; i++ {
		curX := positions[i][0]
		curY := positions[i][1]
		for j := 0; j < m; j++ {
			start := polygens[j][1]
			end := polygens[j][2]
			arr := points[start : end]
			Xs := make([]int, end - start)
			Ys := make([]int, end - start)
			minX := math.MaxInt32
			minY := math.MaxInt32
			maxX := math.MinInt32
			maxY := math.MinInt32
			for k := 0; k < end - start; k++ {
				Xs[k] = arr[k][0]
				Ys[k] = arr[k][1]
				minX = min(minX, Xs[k])
				minY = min(minY, Ys[k])
				maxX = max(maxX, Xs[k])
				maxY = max(maxY, Ys[k])
			}
			if curX < minX || curX > maxX || curY < minY || curY > maxY {
				continue
			}

			//fmt.Println(start - end, curX, curY, Xs, Ys)
			if check(end - start, curX, curY, Xs, Ys) {
				fmt.Println("?")
				res[i] = j + 1
				break
			}
		}
		fmt.Println(res[i])
		if res[i] == 0 {
			res[i] = -1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(pointers int, curX int, curY int, Xs []int, Ys []int) bool {
	i := 0
	j := pointers - 1
	res := false
	for i < pointers {
		if (Ys[i] > curY) != (Ys[j] > curY) && (curX < (Xs[j] - Xs[i]) * (curY - Ys[i]) / (Ys[j] - Ys[i]) + Xs[i]) {
			res = !res
		}
		i++
		j = i
	}
	return res
}


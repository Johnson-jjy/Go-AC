package graph

import "math"

// 解法一: 邻接矩阵 + dijkstra
// 注: 邻接表不好用结点,dijkstra的核心思想:贪心->每次确定已知最短路径节点集合和未知最短路径节点结合

// 解法二: 小根堆 + dijkstra -> 待补充
// 注: 本题主要为稠密图,故解法一更快
func networkDelayTime(times [][]int, n int, k int) int {
	const INF = math.MaxInt32
	store := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		cur := make([]int, n + 1)
		for j := 0; j <= n; j++ {
			cur[j] = INF
		}
		store[i] = cur
	}

	for i := 0; i < len(times); i++ {
		store[times[i][0]][times[i][1]] = times[i][2]
	}

	used := make([]bool, n + 1) // 存已知节点
	dist := make([]int, n + 1) // 存节点的最短路径
	for i := 0; i <= n; i++ {
		dist[i] = INF // 注意初始应该设置为INF,若设置为0,则无法判断路径为0的情况
	}
	dist[k] = 0 // 这样可保证源地为第一个选出的点
	for i := 0; i < n; i++ {
		x := -1
		for k, v := range used {
			if !v && (x == -1 || dist[k] < dist[x]) {
				x = k
			}
		}
		used[x] = true
		for k, d := range store[x] {
			if dist[x] + d < dist[k] {
				dist[k] = dist[x] + d
			}
		}
	}

	res := -1

	for i := 1; i <= n; i++ {
		if dist[i] == INF {
			return -1
		} else if res < dist[i] {
			res = dist[i]
		}
	}

	return res
}

package graph


// 思路:对此刻访问的点做标记(dfs一搜到底,但对点的状态进行记录以便减少爆搜) -> 注意bool可尽早返回
// 区别于以前用的visited bool数组,我们还可以用int的color数组通过数字的不同表示不同的状态 -> 类似三色标记法

// 另: 拓扑排序 -> 待补充
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	res := make([]int, 0)
	color := make([]int, n)

	var safe func(index int) bool
	safe = func(index int) bool {
		if color[index] > 0 {
			return color[index] == 2
		}
		color[index] = 1
		for _, v := range graph[index] {
			if !safe(v) {
				return false
			}
		}

		color[index] = 2
		return true
	}

	for i := 0; i < n; i++ {
		if safe(i) {
			res = append(res, i)
		}
	}

	return res
}

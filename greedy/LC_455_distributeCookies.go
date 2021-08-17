package greedy

import "sort"

// 分发饼干: 排序+贪心->尽可能让每个饼干(饼干不可再分)物尽其用 -> 注意别越界
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0

	index := len(g) - 1
	i := len(s) - 1
	for i >= 0 && index >= 0{
		if s[i] >= g[index] {
			res++
			i--
			index--
		} else {
			index--
		}
	}

	return res
}

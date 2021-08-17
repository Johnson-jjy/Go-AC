package greedy

import "sort"

// 根据身高序列重排: 由高到低或由低到高都可,但实际上由高到低的逻辑是更容易理解的-> 后面低的无论排到哪里都不会影响前面的,只需要按index插入即可
// 注: 1. 关于sort的使用->sort.Slice自己定义排序规则; 2.append做插入时,一定要注意参数的维度
func reconstructQueue(input [][]int) [][]int {
	sort.Slice(input, func(i, j int) bool {
		return input[i][0] > input[j][0] || (input[i][0] == input[j][0] && input[i][1] < input[j][1])
	})

	res := make([][]int, 0)
	for i := 0; i < len(input); i++ {
		cur := input[i]
		index := cur[1]
		// 注意,此处为: append([][]int{cur}
		res = append(res[:index], append([][]int{cur}, res[index:]...)...)
	}

	return res
}

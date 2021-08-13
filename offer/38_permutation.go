package offer

import "sort"

// 字符串排列
// 解法一: 回溯 -> 本质逻辑和去重的全排列一样,但有一些关于byte的注意事项: 1.一定要排序!且对string和bytes无可直接调用的函数,故用sort.Slice
//2. 不能用map!map会导致重复出现的字符在进行used判断时出问题
func permutation(s string) []string {
	res := make([]string, 0)
	used := make([]bool, len(s))
	store := []byte(s)
	// 注: 无排bytes的直接函数
	sort.Slice(store, func(i, j int) bool { return store[i] < store[j] })

	var backtrack func(cur []byte)
	backtrack = func(cur []byte) {
		if len(cur) == len(store) {
			tmp := string(cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(store); i++ {
			if i > 0 && store[i] == store[i - 1] && !used[i - 1] {
				continue
			}
			if !used[i] {
				cur = append(cur, store[i])
				used[i] = true
				backtrack(cur)
				cur = cur[:len(cur) - 1]
				used[i] = false
			}
		}
	}

	cur := make([]byte, 0)
	backtrack(cur)

	return res
}

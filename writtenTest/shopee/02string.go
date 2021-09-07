package main

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 *
 * @param s string字符串
 * @return string字符串一维数组
 */
func permutation(s string) []string {
	res := make([]string, 0)
	used := make([]bool, len(s))
	store := make([]byte, len(s))
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			store[index] = s[i]
			index++
		}
	}
	store = store[:index]
	if len(store) == 0 {
		return res
	}

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


//func permutation(s string) []string {
//	store := []byte(s)
//	sort.Slice(store, func(i, j int) bool { return store[i] < store[j] })
//	n := len(store)
//	perm := make([]byte, 0)
//	visited := make([]bool, n)
//
//	ans := make([]string, 0)
//	var backtrack func(int)
//	backtrack = func(i int) {
//		if i == n {
//			ans = append(ans, string(perm))
//			return
//		}
//
//		for k, v := range visited {
//			if v || k > 0 && !visited[j - 1] && store[j - 1] == store[j] {
//				continue
//			}
//			visited[k] = true
//			perm = append(perm, store[k])
//			backtrack(i + 1)
//			perm = perm[:len(perm) - 1]
//			visited[k] = false
//		}
//	}
//
//	backtrack(0)
//	return ans
//}

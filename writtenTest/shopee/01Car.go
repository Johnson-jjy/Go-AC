package main

func bestFit(V int, item []int) int {
	// write code here
	res := V

	var backtrack func(cur int, index int)
	backtrack = func(cur int, index int) {
		if cur > V {
			return
		}
		if V - cur < res {
			res = V - cur
		}

		for i := index; i < len(item); i++ {
			cur += item[i]
			backtrack(cur, i + 1)
			cur -= item[i]
		}
	}

	backtrack(0, 0)
	return res
}

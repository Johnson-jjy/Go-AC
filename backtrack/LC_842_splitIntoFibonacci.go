package backtrack

import "math"

// 将数组拆分成斐波那契序列: 注意不要用pre12, 回溯的时候不方便修改状态
// 另注: 1.对于getNum, 要验证首字母为0以及得到的数 > MaxInt32的情况
// 2. 对于res的结果,若只有两个数也不能为正确答案
func splitIntoFibonacci(num string) []int {
	res := make([]int, 0)

	getNum := func(start, end int) int {
		if num[start] == '0' && start != end {
			return -1
		}
		sum := 0
		for end >= start {
			cur := int(num[start] - '0')
			sum = sum * 10 + cur
			start++
		}

		return sum
	}

	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		if index == len(num) {
			if len(res) > 2 {
				return true
			}
			return false
		}

		for i := index; i < len(num); i++ {
			cur := getNum(index, i)
			if cur == -1 || cur > math.MaxInt32{
				break
			}
			if len(res) < 2 || cur == res[len(res) - 1] + res[len(res) - 2] {
				res = append(res, cur)
				if backtrack(i + 1) {
					return true
				}
				res = res[:len(res) - 1]
			}
		}

		return false
	}

	if backtrack(0) {
		return res
	}

	return res
}

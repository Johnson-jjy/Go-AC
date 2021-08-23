package string

import (
	"fmt"
	"strconv"
)

// 找出不同的二进制字符串
// 解一: hash存已经有的数字, dfs对每一位上的数字进行试探修改,找到暂无的数字
func findDifferentBinaryString1(nums []string) string {
	n := len(nums)
	store := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		store[nums[i]] = true
	}

	chars := make([]byte, n)
	for i := 0; i < n; i++ {
		chars[i] = '0'
	}

	var res string

	var backtrack func(cur []byte, index int) bool
	backtrack = func(cur []byte, index int) bool {
		if !store[string(cur)] {
			res = string(cur)
			return true
		}

		if index == n {
			return false
		}

		cur[index] = '1'
		if backtrack(cur, index + 1) {
			return true
		}
		cur[index] = '0'

		return false
	}

	backtrack(chars, 0)

	return res
}

// 解二: 本质还是hash + 试探 -> "strconv.ParseInt" -> 直接将字符串转为数字
func findDifferentBinaryString2(nums []string) string {
	has := map[int]bool{}
	for _, s := range nums {
		v, _ := strconv.ParseInt(s, 2, 64) // 用函数可更简洁
		has[int(v)] = true
	}
	v := 0
	for has[v] { v++ } // 找到第一个不存在的
	return fmt.Sprintf("%0*b", len(nums[0]), v) // 按二进制形式打印
}

package backtrack

import (
	"fmt"
	"strconv"
	"strings"
)

// 复原IP地址: 切割类经典 => 可在最后用strings和strconv等进行string的组合
// 注意每次对num的判断; 对每次剩余数字不满足的情况可提前剪枝
// 不建议对原string进行操作,否则回溯时不好处理'.'
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)

	getNum := func(start, end int) int {
		if start != end && (s[start] == '0' || end - start > 2) {
			return -1
		}
		sum := 0
		for start <= end {
			cur := int(s[start] - '0')
			sum = sum * 10 + cur
			start++
		}

		return sum
	}

	var backtrack func(cur []int, index int)
	backtrack = func(cur []int, index int) {
		fmt.Println(cur)
		// 找到了4段且检索完了所有字符,加入res
		if index == len(s) {
			if len(cur) == 4 {
				tmp := make([]string, 4)
				for i := 0; i < 4; i++ {
					tmp[i] = strconv.Itoa(cur[i])
				}
				str := strings.Join(tmp, ".")
				res = append(res, str)
			}
			return
		}

		// 对于剩下数字不可能满足的情况 -> 提前剪枝
		if len(cur) + (len(s) - index + 1)/3 > 4 || len(cur) + (len(s) - index + 1) < 4 {
			return
		}

		for i := index; i < len(s); i++ {
			num := getNum(index, i)
			if num > 255 || num < 0 {
				break
			}
			cur = append(cur, num)
			backtrack(cur, i + 1)
			cur = cur[:len(cur) - 1]
		}
	}

	cur := make([]int, 0)
	backtrack(cur, 0)

	return res
}

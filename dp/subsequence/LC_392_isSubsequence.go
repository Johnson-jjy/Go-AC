package subsequence

// 编辑子序列

/*进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
*/

// 解一: 将其转换为编辑距离类题目 -> 最终判断最长子序列的题目,是否和短串想登
// 1：此处dp数组的定义，表示长度
// 2：对于else--本质是对t的“编辑删除”
func isSubsequence1(s string, t string) bool {
	m, n := len(s), len(t)

	dp := make([][]int, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n + 1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == t[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				dp[i][j] = dp[i][j - 1]
			}
		}
	}

	return dp[m][n] == len(s)
}

// 解二: 直接定义bool数组 -> 注意:初始化1 -> s为空时返回true即可; 2 -> s的第一个字符只要遇到匹配, 就设置为true
// 注! else处 dp[i][j] = dp[i][j - 1] -> 是因为,dp数组的含义是i结尾的s和j结尾的t能否匹配, 必须匹配s中所有字符,故只能删减t
func isSubsequence2(s string, t string) bool {
	m := len(s)
	if m == 0 {
		return true
	}
	n := len(t)
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i - 1] == t[j - 1] {
				if i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i - 1][j - 1]
				}
			} else {
				dp[i][j] = dp[i][j - 1]
			}
		}
	}

	return dp[m][n]
}

// 解3: 双指针移动匹配即可
func isSubsequence3(s string, t string) bool {
	m := len(s)
	if m == 0 {
		return true
	}
	n := len(t)

	left := 0
	right := 0
	for left < m && right < n {
		if s[left] == t[right] {
			left++
		}
		right++
	}

	return left == m
}
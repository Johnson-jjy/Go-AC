package greedy

// 摆动序列: 这种题,上手思路只有两种-> 1.找到某种策略,直接贪心; 2.动归
// 本题解一: 主要是找波峰/波谷的变点;从第一个节点出发,有先向上和先向下两种情况,分别讨论即可
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	flag := false // f找大,t找小
	pre := nums[0]
	count1 := 1
	for i := 1; i < len(nums); i++ {
		if !flag && nums[i] > pre {
			flag = true
			count1++
		} else if flag && nums[i] < pre {
			flag = false
			count1++
		}
		pre = nums[i]
	}

	flag = true
	pre = nums[0]
	count2 := 1
	for i := 1; i < len(nums); i++ {
		if !flag && nums[i] > pre {
			flag = true
			count2++
		} else if flag && nums[i] < pre {
			flag = false
			count2++
		}
		pre = nums[i]
	}

	if count1 > count2 {
		return count1
	}

	return count2
}

// 解法2: 一次遍历; 记录prevDiff和curDiff, 当二者不同号的时候,result++
// 注: curDiff必须严格大于0或小于0,但preDiff可以等于0, 想象一下,一个平的地面可以删除某个节点
func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}

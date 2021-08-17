package greedy

// 加油站: 1.局部最优curSum > 0; 每次 < 0时更新curSum和index;
// 2. 全局最后,sum>0,否则返回-1
// 注: 本质上贪心题目就是两个考虑方向 -> 局部最优 => 全局最优
func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	index := 0
	curSum := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			curSum = 0
			index = i + 1
		}
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		return -1
	}
	return index
}

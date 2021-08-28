package offer

// 数组中出现次数超过一半的数字
// 解一: hash的简单运用
func majorityElement(nums []int) int {
	target := len(nums) / 2 + 1
	store := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		store[nums[i]]++
		if store[nums[i]] == target {
			return nums[i]
		}
	}
	return -1
}

// 解二: 摩尔投票法(待补充)
